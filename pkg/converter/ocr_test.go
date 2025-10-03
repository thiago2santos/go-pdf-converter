package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractTextWithOCR_NonExistentFile(t *testing.T) {
	conv := New(&Config{OCRFallback: true, Verbose: false})

	result, err := conv.extractTextWithOCR("non_existent_file.pdf")

	assert.Error(t, err, "extractTextWithOCR() should return error for non-existent file")
	assert.Empty(t, result, "extractTextWithOCR() should return empty string on error")
}
