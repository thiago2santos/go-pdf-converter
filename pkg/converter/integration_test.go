// +build integration

package converter

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// These tests require actual PDF files in testdata/
// Run with: go test -tags=integration ./...

func TestConvertRealPDF(t *testing.T) {
	// Skip if testdata doesn't exist
	testdataDir := "testdata"
	if _, err := os.Stat(testdataDir); os.IsNotExist(err) {
		t.Skip("testdata directory not found, skipping integration tests")
	}

	// Find PDF files in testdata
	files, err := filepath.Glob(filepath.Join(testdataDir, "*.pdf"))
	if err != nil {
		t.Fatalf("Failed to find test PDFs: %v", err)
	}

	if len(files) == 0 {
		t.Skip("No PDF files in testdata/, skipping integration tests")
	}

	conv := New(DefaultConfig())

	for _, pdfFile := range files {
		t.Run(filepath.Base(pdfFile), func(t *testing.T) {
			result, err := conv.Convert(pdfFile)
			if err != nil {
				t.Errorf("Convert(%s) failed: %v", pdfFile, err)
				return
			}

			// Basic validations
			if result == nil {
				t.Error("Convert() returned nil result")
				return
			}

			if result.Text == "" {
				t.Error("Convert() returned empty text")
			}

			if result.TotalPages == 0 {
				t.Error("Convert() returned 0 pages")
			}

			if result.Method == "" {
				t.Error("Convert() returned empty method")
			}

			// Check statistics are calculated
			if result.CharactersCount == 0 {
				t.Error("Convert() returned 0 character count")
			}

			t.Logf("Converted %s: %d pages, %d words, method: %s",
				filepath.Base(pdfFile),
				result.TotalPages,
				result.WordsCount,
				result.Method)
		})
	}
}

func TestConvertWithOCRFallback(t *testing.T) {
	testdataDir := "testdata"
	if _, err := os.Stat(testdataDir); os.IsNotExist(err) {
		t.Skip("testdata directory not found, skipping integration tests")
	}

	files, err := filepath.Glob(filepath.Join(testdataDir, "*.pdf"))
	if err != nil || len(files) == 0 {
		t.Skip("No PDF files in testdata/, skipping integration tests")
	}

	// Test with OCR enabled
	conv := New(&Config{
		OCRFallback: true,
		Verbose:     false,
	})

	result, err := conv.Convert(files[0])
	if err != nil {
		t.Fatalf("Convert() with OCR fallback failed: %v", err)
	}

	if result == nil {
		t.Fatal("Convert() returned nil result")
	}

	// Should have extracted some content
	if len(strings.TrimSpace(result.Text)) == 0 {
		t.Error("Convert() extracted no text")
	}
}

func TestConvertWithoutOCRFallback(t *testing.T) {
	testdataDir := "testdata"
	if _, err := os.Stat(testdataDir); os.IsNotExist(err) {
		t.Skip("testdata directory not found, skipping integration tests")
	}

	files, err := filepath.Glob(filepath.Join(testdataDir, "*.pdf"))
	if err != nil || len(files) == 0 {
		t.Skip("No PDF files in testdata/, skipping integration tests")
	}

	// Test with OCR disabled
	conv := New(&Config{
		OCRFallback: false,
		Verbose:     false,
	})

	result, err := conv.Convert(files[0])
	if err != nil {
		t.Fatalf("Convert() without OCR fallback failed: %v", err)
	}

	if result == nil {
		t.Fatal("Convert() returned nil result")
	}

	// Method should be text extraction only
	if result.Method == MethodOCR {
		t.Error("Convert() used OCR when OCRFallback was disabled")
	}
}

