package tests

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// minResourceMethods guards the scan against passing vacuously (e.g. wrong directory).
const minResourceMethods = 70

// TestResourceMethodReturnTypesMatchDeserializedResponse parses every resource file in the
// SDK root and asserts that each exported resource method's declared *operations.X result is
// the same type it (or the parse helper it delegates to) actually constructs on the success path.
func TestResourceMethodReturnTypesMatchDeserializedResponse(t *testing.T) {
	root, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, root, func(fi os.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg, ok := pkgs["fastpixgo"]
	if !ok {
		t.Fatalf("package fastpixgo not found in %s", root)
	}

	// name -> declared *operations.X result, and name -> response type literal built in the body.
	type fn struct {
		declared string
		built    string
		calls    []string
		exported bool
		receiver string
	}
	fns := map[string]*fn{}
	for _, file := range pkg.Files {
		for _, d := range file.Decls {
			fd, ok := d.(*ast.FuncDecl)
			if !ok || fd.Recv == nil || fd.Body == nil {
				continue
			}
			declared := operationsResult(fd.Type)
			if declared == "" {
				continue
			}
			f := &fn{declared: declared, exported: fd.Name.IsExported(), receiver: receiverName(fd.Recv)}
			ast.Inspect(fd.Body, func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.UnaryExpr:
					if x.Op == token.AND {
						if cl, ok := x.X.(*ast.CompositeLit); ok {
							if name := operationsType(cl.Type); name != "" && strings.HasSuffix(name, "Response") {
								f.built = name
							}
						}
					}
				case *ast.CallExpr:
					if sel, ok := x.Fun.(*ast.SelectorExpr); ok {
						f.calls = append(f.calls, sel.Sel.Name)
					}
				}
				return true
			})
			fns[f.receiver+"."+fd.Name.Name] = f
		}
	}

	resolveBuilt := func(f *fn) string {
		if f.built != "" {
			return f.built
		}
		for _, c := range f.calls {
			if strings.HasPrefix(c, "parse") {
				if h, ok := fns[f.receiver+"."+c]; ok && h.built != "" {
					return h.built
				}
			}
		}
		return ""
	}

	checked := 0
	for name, f := range fns {
		if !f.exported {
			continue
		}
		built := resolveBuilt(f)
		if built == "" {
			t.Errorf("%s: could not find the *operations.*Response literal it builds (extend the scanner, do not skip)", name)
			continue
		}
		if built != f.declared {
			t.Errorf("%s declares *operations.%s but builds *operations.%s", name, f.declared, built)
		}
		checked++
	}
	if checked < minResourceMethods {
		t.Fatalf("scanned only %d resource methods, expected at least %d", checked, minResourceMethods)
	}
	t.Logf("checked %d resource methods", checked)
}

func receiverName(recv *ast.FieldList) string {
	if recv == nil || len(recv.List) == 0 {
		return ""
	}
	if star, ok := recv.List[0].Type.(*ast.StarExpr); ok {
		if id, ok := star.X.(*ast.Ident); ok {
			return id.Name
		}
	}
	return ""
}

// operationsResult returns X when the first result is *operations.X.
func operationsResult(ft *ast.FuncType) string {
	if ft.Results == nil || len(ft.Results.List) == 0 {
		return ""
	}
	star, ok := ft.Results.List[0].Type.(*ast.StarExpr)
	if !ok {
		return ""
	}
	return operationsType(star.X)
}

func operationsType(e ast.Expr) string {
	sel, ok := e.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	if pkg, ok := sel.X.(*ast.Ident); ok && pkg.Name == "operations" {
		return sel.Sel.Name
	}
	return ""
}
