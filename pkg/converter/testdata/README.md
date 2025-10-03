# Test Data

This directory contains PDF files used for integration testing.

## Adding Test PDFs

To run integration tests, add sample PDF files to this directory:

```bash
# Add your test PDFs here
cp ~/sample.pdf testdata/
```

## Running Tests

**Unit tests only (default):**
```bash
go test ./...
```

**Integration tests (requires PDF files):**
```bash
go test -tags=integration ./...
```

**All tests with coverage:**
```bash
go test -tags=integration -coverprofile=coverage.txt ./...
```

## Test PDF Types

Consider adding various PDF types:
- Text-based PDFs (extractable text)
- Image-based PDFs (requires OCR)
- Mixed content PDFs
- Multi-page documents
- Different encodings

## Note

Test PDF files are ignored by git (see `.gitignore`) to keep the repository size small.
Add your own test files locally for comprehensive testing.

