//go:build !cgo
// +build !cgo

package converter

import (
	"fmt"
)

// extractTextWithOCR is a stub for non-CGO builds.
func (c *Converter) extractTextWithOCR(pdfPath string) (string, error) {
	return "", fmt.Errorf("OCR is not available in this build. To enable OCR, build from source with CGO_ENABLED=1 and Tesseract installed")
}

