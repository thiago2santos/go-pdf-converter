# Contributing to go-pdf-converter

First off, thank you for considering contributing to go-pdf-converter! It's people like you that make it a great tool.

## Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates. When you create a bug report, include as many details as possible:

* **Use a clear and descriptive title**
* **Describe the exact steps to reproduce the problem**
* **Provide specific examples** (sample PDF files if possible)
* **Describe the behavior you observed** and what you expected
* **Include your environment details** (OS, Go version, Tesseract version)

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion:

* **Use a clear and descriptive title**
* **Provide a detailed description of the suggested enhancement**
* **Explain why this enhancement would be useful**
* **List any alternatives you've considered**

### Pull Requests

1. Fork the repo and create your branch from `main`
2. If you've added code, add tests
3. Ensure the test suite passes
4. Make sure your code follows the existing style
5. Write a clear commit message
6. Update documentation as needed

## Development Setup

### Prerequisites

```bash
# Install Go 1.21 or later
go version

# Install Tesseract OCR
# macOS
brew install tesseract

# Ubuntu/Debian
sudo apt-get install tesseract-ocr libtesseract-dev

# Fedora/RHEL
sudo dnf install tesseract tesseract-devel
```

### Building from Source

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/go-pdf-converter.git
cd go-pdf-converter

# Install dependencies
go mod download

# Install git hooks (recommended)
make install-hooks

# Build
make build

# Run tests
make test

# Run linter
make lint
```

### Git Hooks

We use pre-commit hooks to ensure code quality. The hooks will:
- Check code formatting with `gofmt`
- Run `golangci-lint` before each commit

**Install hooks:**
```bash
make install-hooks
```

**Skip hooks (when needed):**
```bash
git commit --no-verify
```

### Project Structure

```
.
├── cmd/
│   └── pdf-converter/      # CLI application
├── pkg/
│   └── converter/          # Core library (reusable)
├── examples/               # Example usage
├── docs/                   # Documentation
└── .github/               # GitHub specific files
```

## Coding Guidelines

### Go Style

* Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines
* Use `gofmt` to format your code
* Run `golangci-lint` before submitting
* Write meaningful commit messages
* Add comments for exported functions and types
* Keep functions small and focused

### Testing

* Write tests for new functionality
* Maintain or improve code coverage
* Include both unit and integration tests
* Test edge cases and error conditions

### Documentation

* Update README.md if you change functionality
* Add godoc comments for exported items
* Include examples for new features
* Keep documentation clear and concise

## Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters
* Reference issues and pull requests when relevant

Examples:
```
Add OCR language selection option

Implement support for selecting OCR language through CLI flag.
Closes #123

Fix memory leak in PDF processing

Properly close file handles after processing each page.
Fixes #456
```

## Review Process

1. Maintainers will review your PR
2. Address any feedback or requested changes
3. Once approved, your PR will be merged
4. Your contribution will be included in the next release

## Questions?

Feel free to open an issue with your question or reach out to the maintainers.

Thank you for contributing! 🎉

