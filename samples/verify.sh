#!/usr/bin/env bash
#
# Compile-checks every example in this directory against the local SDK.
#
# The samples carry a `//go:build ignore` constraint (so `go build ./...` skips
# them) and each declares its own `package main`/`func main()`. To type-check
# them we strip the constraint and build each one in an isolated throwaway
# module that points back at this checkout via a replace directive.
#
# Exits non-zero if any sample fails to compile. Run from anywhere.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MOD_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

fail=0
for sample in "$SCRIPT_DIR"/*.go; do
  base="$(basename "$sample" .go)"
  dir="$WORK/$base"
  mkdir -p "$dir"
  grep -vE '^//go:build ignore$|^// \+build ignore$' "$sample" > "$dir/main.go"
  cat > "$dir/go.mod" <<EOF
module samplecheck/$base
go 1.22
require github.com/FastPix/fastpix-go v0.0.0
replace github.com/FastPix/fastpix-go => $MOD_ROOT
EOF
  ( cd "$dir" && go mod tidy >/dev/null 2>&1 || true )
  if out="$(cd "$dir" && go build ./... 2>&1)"; then
    echo "OK   $base"
  else
    echo "FAIL $base"
    echo "$out"
    fail=1
  fi
done

exit $fail
