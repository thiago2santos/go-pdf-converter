//go:build integration
// +build integration

package converter

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type integrationTestData struct {
	pdfFile     string
	config      *Config
	expectError bool
}

func TestIntegration_ConvertRealPDF(t *testing.T) {
	testdataDir := "testdata"
	if _, err := os.Stat(testdataDir); os.IsNotExist(err) {
		t.Skip("testdata directory not found, skipping integration tests")
	}

	files, err := filepath.Glob(filepath.Join(testdataDir, "*.pdf"))
	require.NoError(t, err, "Failed to find test PDFs")

	if len(files) == 0 {
		t.Skip("No PDF files in testdata/, skipping integration tests")
	}

	conv := New(DefaultConfig())

	for _, pdfFile := range files {
		t.Run(filepath.Base(pdfFile), func(t *testing.T) {
			result, err := conv.Convert(pdfFile)

			require.NoError(t, err, "Convert(%s) should not fail", pdfFile)
			require.NotNil(t, result, "Convert() should return non-nil result")

			assert.NotEmpty(t, result.Text, "Convert() should extract text")
			assert.NotZero(t, result.TotalPages, "Convert() should have at least one page")
			assert.NotEmpty(t, result.Method, "Convert() should specify extraction method")
			assert.NotZero(t, result.CharactersCount, "Convert() should count characters")

			t.Logf("Converted %s: %d pages, %d words, method: %s",
				filepath.Base(pdfFile),
				result.TotalPages,
				result.WordsCount,
				result.Method)
		})
	}
}

func TestIntegration_ConvertWithOCRConfig(t *testing.T) {
	testdataDir := "testdata"
	if _, err := os.Stat(testdataDir); os.IsNotExist(err) {
		t.Skip("testdata directory not found, skipping integration tests")
	}

	files, err := filepath.Glob(filepath.Join(testdataDir, "*.pdf"))
	if err != nil || len(files) == 0 {
		t.Skip("No PDF files in testdata/, skipping integration tests")
	}

	tests := []struct {
		name      string
		config    *Config
		expectOCR bool
	}{
		{
			name: "Success - with OCR fallback enabled",
			config: &Config{
				OCRFallback: true,
				Verbose:     false,
			},
			expectOCR: false,
		},
		{
			name: "Success - with OCR fallback disabled",
			config: &Config{
				OCRFallback: false,
				Verbose:     false,
			},
			expectOCR: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conv := New(tt.config)
			result, err := conv.Convert(files[0])

			require.NoError(t, err, "Convert() should not fail")
			require.NotNil(t, result, "Convert() should return non-nil result")
			assert.NotEmpty(t, strings.TrimSpace(result.Text), "Convert() should extract text")

			if !tt.config.OCRFallback {
				assert.NotEqual(t, MethodOCR, result.Method,
					"Convert() should not use OCR when OCRFallback is disabled")
			}
		})
	}
}
