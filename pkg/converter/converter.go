// Package converter provides PDF text extraction capabilities with OCR fallback.
package converter

import (
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

// ExtractionMethod represents the method used to extract text from a PDF.
type ExtractionMethod string

const (
	// MethodTextExtraction indicates direct text extraction was used.
	MethodTextExtraction ExtractionMethod = "Text Extraction"
	// MethodOCR indicates OCR was used to extract text.
	MethodOCR ExtractionMethod = "OCR"
)

// Result contains the extracted text and metadata about the conversion.
type Result struct {
	Text            string
	Method          ExtractionMethod
	TotalPages      int
	LinesCount      int
	WordsCount      int
	CharactersCount int
}

// Config holds configuration options for PDF conversion.
type Config struct {
	// OCRFallback enables OCR when no text is found
	OCRFallback bool
	// Verbose enables detailed logging
	Verbose bool
}

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		OCRFallback: true,
		Verbose:     false,
	}
}

// Converter handles PDF to text conversion.
type Converter struct {
	config *Config
}

// New creates a new Converter with the given configuration.
func New(config *Config) *Converter {
	if config == nil {
		config = DefaultConfig()
	}
	return &Converter{config: config}
}

// Convert converts a PDF file to text, using OCR if necessary.
func (c *Converter) Convert(pdfPath string) (*Result, error) {
	// Try regular text extraction first
	text, totalPages, err := c.extractTextFromPDF(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("failed to extract text: %w", err)
	}

	method := MethodTextExtraction
	textContent := strings.TrimSpace(text)

	// If no meaningful text found and OCR is enabled, try OCR
	if c.config.OCRFallback && (textContent == "" || strings.Contains(textContent, "[No text content found on this page]")) {
		ocrText, err := c.extractTextWithOCR(pdfPath)
		if err != nil {
			text += fmt.Sprintf("\nWARNING: OCR failed: %v\n", err)
		} else {
			text = ocrText
			method = MethodOCR
		}
	}

	// Calculate statistics
	lines := len(strings.Split(text, "\n"))
	words := len(strings.Fields(text))
	chars := len(text)

	return &Result{
		Text:            text,
		Method:          method,
		TotalPages:      totalPages,
		LinesCount:      lines,
		WordsCount:      words,
		CharactersCount: chars,
	}, nil
}

// extractTextFromPDF extracts text directly from a PDF file.
func (c *Converter) extractTextFromPDF(pdfPath string) (result string, pages int, err error) {
	f, r, err := pdf.Open(pdfPath)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()

	var text strings.Builder
	pages = r.NumPage()

	text.WriteString(fmt.Sprintf("Total Pages: %d\n\n", pages))

	for pageIndex := 1; pageIndex <= pages; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		text.WriteString(fmt.Sprintf("=== PAGE %d ===\n", pageIndex))

		pageText, err := p.GetPlainText(nil)
		if err != nil {
			text.WriteString(fmt.Sprintf("[Error reading page %d: %v]\n", pageIndex, err))
			continue
		}

		if strings.TrimSpace(pageText) != "" {
			text.WriteString(pageText)
		} else {
			text.WriteString("[No text content found on this page]\n")
		}
		text.WriteString("\n\n")
	}

	return text.String(), pages, nil
}
