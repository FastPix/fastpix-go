---
name: Bug Report
about: Report a bug or unexpected behavior in the FastPix Go SDK
title: '[BUG] '
labels: ['bug', 'needs-triage']
assignees: ''
---

# Bug Report

Thank you for taking the time to report a bug with the FastPix Go SDK. To help us resolve your issue quickly and efficiently, please provide the following information:

## Description
**Clear and concise description of the bug:**
```
<!-- Please provide a detailed description of what you're experiencing -->
```

## Environment Information

### System Details
- **Go Version:** [e.g., 1.22, 1.23, 1.24]
- **Operating System:** [e.g., Windows 10, macOS 12.0, Ubuntu 20.04, etc.]
- **Architecture:** [e.g., amd64, arm64, etc.] (Optional but helpful)

### SDK Information
- **FastPix Go SDK Version:** [e.g., 1.0.3, 1.0.2, etc.]
- **Go Modules:** [Enabled/Disabled]

## Reproduction Steps

1. **Setup Environment:**
   ```bash
   go get github.com/fastpix/fastpix-go
   ```

2. **Code to Reproduce:**
   ```go
   // Please provide a minimal, reproducible example
   package main

   import (
       "context"
       "github.com/fastpix/fastpix-go"
   )

   func main() {
       ctx := context.Background()
       fastpix := fastpixgo.New(
           fastpixgo.WithSecurity(fastpixgo.Security{
               Username: "your-username",
               Password: "your-password",
           }),
       )
       
       // Your code here that causes the issue
   }
   ```

3. **Expected Behavior:**

    ```
    <!-- Describe what you expected to happen -->
    ```

4. **Actual Behavior:**

    ```
    <!-- Describe what actually happened -->
    ```

5. **Error Messages/Logs:**
   ```
   <!-- Paste any error messages, stack traces, or logs here -->
   ```

## Debugging Information

### Console Output
```
<!-- Paste the complete console output here -->
```

### Error Stack Traces
```go
// Complete stack trace for Go errors
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1234567]

goroutine 1 [running]:
github.com/fastpix/fastpix-go.SomeMethod(...)
    /path/to/file.go:123
main.main()
    /path/to/your/file.go:45
```

### HTTP Requests
```http
# Raw HTTP request (remove sensitive headers and credentials)
POST /api/endpoint HTTP/1.1
Host: [FastPix API endpoint]
Authorization: Basic ***
Content-Type: application/json

<!-- Remove credentials and sensitive headers before pasting -->
```

### Screenshots
```
<!-- If applicable, please attach screenshots that help explain your issue -->
```

## Additional Context

### Configuration
```go
// Please share your SDK configuration (remove sensitive information)
fastpix := fastpixgo.New(
    fastpixgo.WithSecurity(fastpixgo.Security{
        Username: "***", // Redacted
        Password: "***", // Redacted
    }),
    // Any other configuration options
)
```

### Workarounds
```
<!-- If you've found any workarounds, please describe them here -->
```

## Priority
Please indicate the priority of this bug:

- [ ] Critical (Blocks production use)
- [ ] High (Significant impact on functionality)
- [ ] Medium (Minor impact)
- [ ] Low (Nice to have)

## Checklist
Before submitting, please ensure:

- [ ] I have searched existing issues to avoid duplicates
- [ ] I have provided all required information
- [ ] I have tested with the latest SDK version
- [ ] I have removed any sensitive information (credentials, API keys, etc.)
- [ ] I have provided a minimal reproduction case
- [ ] I have checked the documentation

---

**Thank you for helping improve the FastPix Go SDK! 🚀**

