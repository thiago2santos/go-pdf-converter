# 🔄 go-pdf-converter

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/thiago2santos/go-pdf-converter)](https://goreportcard.com/report/github.com/thiago2santos/go-pdf-converter)
[![codecov](https://codecov.io/gh/thiago2santos/go-pdf-converter/branch/main/graph/badge.svg)](https://codecov.io/gh/thiago2santos/go-pdf-converter)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

A fast, efficient PDF to text converter with automatic OCR fallback. Extract text from any PDF—whether it's text-based or image-based.

## ✨ Features

- 🚀 **Blazing Fast**: Direct text extraction for text-based PDFs (~0.1s)
- 🤖 **Smart OCR Fallback**: Automatically detects and processes image-based PDFs
- 📦 **Single Binary**: No runtime dependencies except Tesseract for OCR
- 🎯 **Simple API**: Use as CLI tool or Go library
- 🌍 **Cross-Platform**: Works on macOS, Linux, and Windows
- 📊 **Rich Output**: Detailed statistics and formatted text with page markers
- 🧪 **Battle-Tested**: Validated on 29+ different PDF types

## 📥 Installation

### Prerequisites

**For OCR support (image-based PDFs), install Tesseract:**

```bash
# macOS
brew install tesseract

# Ubuntu/Debian
sudo apt-get install tesseract-ocr libtesseract-dev

# Fedora/RHEL
sudo dnf install tesseract tesseract-devel

# Windows (via Chocolatey)
choco install tesseract
```

### Option 1: Install Pre-built Binary (Recommended)

Download the latest release for your platform from the [releases page](https://github.com/thiago2santos/go-pdf-converter/releases).

```bash
# Extract and move to your PATH
tar -xzf pdf-converter_*.tar.gz
sudo mv pdf-converter /usr/local/bin/
```

**Note:** Pre-built binaries support text extraction but **not OCR**. For OCR support, use Option 3 (build from source with Tesseract).

### Option 2: Install with Go

```bash
go install github.com/thiago2santos/go-pdf-converter/cmd/pdf-converter@latest
```

### Option 3: Build from Source (with OCR support)

To enable OCR support, first install Tesseract (see Prerequisites above), then:

```bash
git clone https://github.com/thiago2santos/go-pdf-converter.git
cd go-pdf-converter

# Build with OCR support (requires Tesseract)
make build

# Binary will be in ./bin/pdf-converter with full OCR support
# Optional: Install to PATH
sudo cp bin/pdf-converter /usr/local/bin/
```

## 🚀 Quick Start

### CLI Usage

```bash
# Convert a PDF to text
pdf-converter document.pdf

# Output is saved as document_extracted.txt
```

### As a Go Library

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/thiago2santos/go-pdf-converter/pkg/converter"
)

func main() {
    // Create a new converter
    config := converter.DefaultConfig()
    config.OCRFallback = true // Enable OCR for image PDFs
    
    conv := converter.New(config)
    
    // Convert PDF
    result, err := conv.Convert("document.pdf")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Extracted %d words using %s\n", 
        result.WordsCount, result.Method)
    fmt.Println(result.Text)
}
```

## 📖 Examples

### Converting Text-Based PDFs

```bash
$ pdf-converter invoice.pdf
🔄 Converting 'invoice.pdf'...
✅ Conversion completed using Text Extraction
📄 Text saved to: invoice_extracted.txt
📊 Stats: 3 pages, 245 lines, 1,847 words, 12,453 characters
```

### Converting Image-Based PDFs (OCR)

```bash
$ pdf-converter scanned-document.pdf
🔄 Converting 'scanned-document.pdf'...
📄 No extractable text found, trying OCR...
  Processing page 1/5 with OCR...
  Processing page 2/5 with OCR...
  ...
✅ Conversion completed using OCR
📄 Text saved to: scanned-document_extracted.txt
📊 Stats: 5 pages, 423 lines, 3,124 words, 19,876 characters
```

## 📝 Output Format

```
PDF: document.pdf
Extraction Method: Text Extraction

Total Pages: 3

=== PAGE 1 ===
[Extracted text content...]

=== PAGE 2 ===
[More content...]

=== PAGE 3 ===
[Final page content...]
```

## ⚡ Performance

| PDF Type | Method | Speed | Accuracy |
|----------|--------|-------|----------|
| Text-based | Direct extraction | ~0.1s | 100% |
| Image-based | OCR (Tesseract) | ~2-10s/page | 95%+ |

## 🏗️ Architecture

```
┌─────────────────┐
│  PDF Document   │
└────────┬────────┘
         │
         v
┌─────────────────────────────┐
│  Try Text Extraction First  │ ← Fast path
└────────┬────────────────────┘
         │
         │ No text found?
         v
┌─────────────────────────────┐
│  Convert Pages to Images    │
└────────┬────────────────────┘
         │
         v
┌─────────────────────────────┐
│  OCR with Tesseract         │ ← Fallback
└────────┬────────────────────┘
         │
         v
┌─────────────────────────────┐
│  Formatted Text Output      │
└─────────────────────────────┘
```

## 🛠️ Development

### Running Tests

```bash
# Run unit tests
make test

# Run tests with race detector (requires CGO)
make test-race

# Run integration tests (requires test PDFs in testdata/)
make test-integration

# Run tests with coverage
make coverage
```

### Test Coverage

We maintain comprehensive unit tests for the converter library:
- Configuration and initialization
- Error handling
- Input validation
- Method selection logic

Add test PDFs to `pkg/converter/testdata/` for integration testing.

### Running Linter

```bash
make lint
```

### Code Quality

```bash
# Format code
make fmt

# Run all checks (fmt, lint, test)
make check

# Install pre-commit hooks
make install-hooks
```

### Building

```bash
make build
```

### See All Commands

```bash
make help
```

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) first.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [ledongthuc/pdf](https://github.com/ledongthuc/pdf) - PDF text extraction
- [gen2brain/go-fitz](https://github.com/gen2brain/go-fitz) - PDF to image conversion
- [otiai10/gosseract](https://github.com/otiai10/gosseract) - Tesseract OCR bindings
- [Tesseract OCR](https://github.com/tesseract-ocr/tesseract) - OCR engine

## 📬 Contact

- GitHub: [@thiago2santos](https://github.com/thiago2santos)
- Issues: [Report a bug](https://github.com/thiago2santos/go-pdf-converter/issues)

## ⭐ Show Your Support

If you find this project useful, please give it a star! It helps others discover the project.

---

**Made with ❤️ by the open-source community**
