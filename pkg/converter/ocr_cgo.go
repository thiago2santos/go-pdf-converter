//go:build cgo
// +build cgo

package converter

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"github.com/gen2brain/go-fitz"
	"github.com/otiai10/gosseract/v2"
)

// extractTextWithOCR extracts text from a PDF using OCR.
func (c *Converter) extractTextWithOCR(pdfPath string) (string, error) {
	// Open PDF with go-fitz for image conversion
	doc, err := fitz.New(pdfPath)
	if err != nil {
		return "", fmt.Errorf("failed to open PDF: %w", err)
	}
	defer doc.Close()

	// Initialize OCR client
	ocrClient := gosseract.NewClient()
	defer ocrClient.Close()

	var text strings.Builder
	var tmpFiles []string
	defer func() {
		for _, tmpFile := range tmpFiles {
			os.Remove(tmpFile)
		}
	}()

	text.WriteString(fmt.Sprintf("Total Pages: %d\n\n", doc.NumPage()))

	for i := 0; i < doc.NumPage(); i++ {
		if c.config.Verbose {
			fmt.Printf("  Processing page %d/%d with OCR...\n", i+1, doc.NumPage())
		}

		text.WriteString(fmt.Sprintf("=== PAGE %d ===\n", i+1))

		// Convert page to image
		img, err := doc.Image(i)
		if err != nil {
			text.WriteString(fmt.Sprintf("[Error converting page %d to image: %v]\n", i+1, err))
			continue
		}

		// Create temporary file for the image
		tmpFile, err := os.CreateTemp("", "page_*.png")
		if err != nil {
			text.WriteString(fmt.Sprintf("[Error creating temp file for page %d: %v]\n", i+1, err))
			continue
		}
		tmpFileName := tmpFile.Name()

		// Save image to temp file
		err = png.Encode(tmpFile, img)
		tmpFile.Close()
		if err != nil {
			os.Remove(tmpFileName)
			text.WriteString(fmt.Sprintf("[Error saving image for page %d: %v]\n", i+1, err))
			continue
		}
		tmpFiles = append(tmpFiles, tmpFileName)

		// Perform OCR on the image
		err = ocrClient.SetImage(tmpFileName)
		if err != nil {
			text.WriteString(fmt.Sprintf("[Error setting OCR image for page %d: %v]\n", i+1, err))
			continue
		}

		pageText, err := ocrClient.Text()
		if err != nil {
			text.WriteString(fmt.Sprintf("[OCR error on page %d: %v]\n", i+1, err))
			continue
		}

		if strings.TrimSpace(pageText) != "" {
			text.WriteString(pageText)
		} else {
			text.WriteString("[No text detected via OCR]\n")
		}
		text.WriteString("\n\n")
	}

	return text.String(), nil
}
