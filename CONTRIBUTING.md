# Contributing to vexErrScope

Thank you for your interest in contributing to vexErrScope! This document provides guidelines and instructions for contributing.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/vexErrScope.git`
3. Create a branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -m "Add your feature"`
7. Push to your fork: `git push origin feature/your-feature-name`
8. Open a Pull Request

## Development Setup

```bash
# Clone the repository
git clone https://github.com/vextmdev/vexErrScope.git
cd vexErrScope

# Run tests
go test ./...

# Build
go build -o vexErrScope ./cmd/vexErrScope

# Run
./vexErrScope explain testdata/sample.log
```

## Code Guidelines

### Style

- Follow standard Go conventions and `gofmt`
- Use meaningful variable and function names
- Keep functions focused and small
- Add comments only where necessary to explain "why", not "what"

### Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting a PR
- Add test cases for edge cases and error conditions

### Commits

- Use clear, descriptive commit messages
- Keep commits focused on a single change
- Reference issues in commit messages when applicable (e.g., "Fix #123")

## Adding New Error Patterns

To add support for a new error pattern:

1. Add pattern detection in `internal/parse/parser.go`
2. Add pattern mapping in `internal/analyze/analyzer.go`
3. Add corresponding unit tests
4. Update documentation

Example pattern structure in `analyzer.go`:

```go
"your_pattern_name": {
    ErrorType:  "Human Readable Name",
    RootCause:  "Description of why this error occurs.",
    SuggestFix: "Steps to fix this error.",
    Confidence: "high|medium|low",
},
```

## Pull Request Process

1. Ensure your code follows the project's style guidelines
2. Update documentation if needed
3. Add or update tests as appropriate
4. Ensure all tests pass
5. Request review from maintainers

## Reporting Bugs

When reporting bugs, please include:

- Go version (`go version`)
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
- Sample input that triggers the bug (if applicable)

## Feature Requests

Feature requests are welcome! Please open an issue describing:

- The problem you're trying to solve
- Your proposed solution
- Any alternatives you've considered

## Code of Conduct

Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md).

## Questions?

Feel free to open an issue for any questions about contributing.
