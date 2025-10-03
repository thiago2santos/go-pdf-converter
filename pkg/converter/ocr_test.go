package converter

import (
	"testing"
)

func TestExtractTextWithOCRNonExistentFile(t *testing.T) {
	conv := New(&Config{OCRFallback: true, Verbose: false})

	_, err := conv.extractTextWithOCR("non_existent_file.pdf")

	if err == nil {
		t.Error("extractTextWithOCR() with non-existent file should return error")
	}
}

func TestOCRFallbackDisabled(t *testing.T) {
	conv := New(&Config{OCRFallback: false, Verbose: false})

	if conv.config.OCRFallback {
		t.Error("OCRFallback should be disabled")
	}
}

