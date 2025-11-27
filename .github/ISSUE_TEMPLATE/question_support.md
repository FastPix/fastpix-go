---
name: Question/Support
about: Ask questions or get help with the FastPix Go SDK
title: '[QUESTION] '
labels: ['question', 'needs-triage']
assignees: ''
---

# Question/Support

Thank you for reaching out! We're here to help you with the FastPix Go SDK. Please provide the following information:

## Question Type
- [ ] How to use a specific feature
- [ ] Integration help
- [ ] Configuration question
- [ ] Performance question
- [ ] Troubleshooting help
- [ ] Context handling
- [ ] Error handling
- [ ] Goroutines/concurrency
- [ ] Other: _______________

## Question
**What would you like to know?**
```
<!-- Please provide a clear, specific question -->
```
## What You've Tried
**What have you already attempted to solve this?**

```go
// Please share any code you've tried
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
    
    // Your attempted code here
}
```

## Current Setup
**Describe your current setup:**

### Environment
- **Go Version:** [e.g., 1.22, 1.23, 1.24]
- **Operating System:** [e.g., Windows 10, macOS 12.0, Ubuntu 20.04, etc.]
- **FastPix Go SDK Version:** [e.g., 1.0.3, 1.0.2]
- **Architecture:** [e.g., amd64, arm64, etc.]

### Configuration
```go
// Your current SDK configuration (remove sensitive information)
fastpix := fastpixgo.New(
    fastpixgo.WithSecurity(fastpixgo.Security{
        Username: "***", // Redacted
        Password: "***", // Redacted
    }),
    // Any other configuration
)
```

## Expected Outcome
**What are you trying to achieve?**
```
<!-- Describe your end goal -->
```
## Error Messages (if any)
```
<!-- If you're getting errors, paste them here -->
```

## Additional Context

### Use Case
**What are you building?**

- [ ] Web service/API
- [ ] CLI application
- [ ] Microservice
- [ ] Background worker
- [ ] Library/package
- [ ] Other: _______________

### Project Details
- **Project Type:** [e.g., REST API, gRPC service, CLI tool, etc.]
- **Go Modules:** [Enabled/Disabled]

### Timeline
**When do you need this resolved?**

- [ ] ASAP (blocking development)
- [ ] This week
- [ ] This month
- [ ] No rush

### Resources Checked
**What resources have you already checked?**

- [ ] README.md
- [ ] Documentation
- [ ] Examples
- [ ] Stack Overflow
- [ ] GitHub Issues
- [ ] Go documentation
- [ ] Other: _______________

## Priority
Please indicate the urgency:

- [ ] Critical (Blocking production deployment)
- [ ] High (Blocking development)
- [ ] Medium (Would like to know soon)
- [ ] Low (Just curious)

## Checklist
Before submitting, please ensure:

- [ ] I have provided a clear question
- [ ] I have described what I've tried
- [ ] I have included my current setup
- [ ] I have checked existing documentation
- [ ] I have provided sufficient context
- [ ] I have removed any sensitive information (credentials, API keys, etc.)

---

**We'll do our best to help you get unstuck! 🚀**

**For urgent issues, please also consider:**
- [FastPix Documentation](https://docs.fastpix.io/)
- [Stack Overflow](https://stackoverflow.com/questions/tagged/fastpix)
- [GitHub Discussions](https://github.com/FastPix/golang-server-side-language-sdk/discussions)

