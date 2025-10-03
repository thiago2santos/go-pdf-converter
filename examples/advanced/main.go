package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thiago2santos/go-pdf-converter/pkg/converter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <pdf_file>")
		os.Exit(1)
	}

	pdfPath := os.Args[1]

	// Create custom configuration
	config := &converter.Config{
		OCRFallback: true, // Enable OCR fallback
		Verbose:     true, // Enable verbose logging
	}

	// Create converter
	conv := converter.New(config)

	// Convert PDF
	fmt.Println("Starting conversion...")
	result, err := conv.Convert(pdfPath)
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	}

	// Analyze results
	fmt.Printf("\n" + "="*50 + "\n")
	fmt.Printf("CONVERSION SUMMARY\n")
	fmt.Printf("="*50 + "\n")
	fmt.Printf("File: %s\n", pdfPath)
	fmt.Printf("Method: %s\n", result.Method)
	fmt.Printf("Pages: %d\n", result.TotalPages)
	fmt.Printf("Lines: %d\n", result.LinesCount)
	fmt.Printf("Words: %d\n", result.WordsCount)
	fmt.Printf("Characters: %d\n", result.CharactersCount)

	// Calculate average words per page
	if result.TotalPages > 0 {
		avgWords := float64(result.WordsCount) / float64(result.TotalPages)
		fmt.Printf("Avg words/page: %.2f\n", avgWords)
	}

	// Save to file
	outputFile := "output.txt"
	err = os.WriteFile(outputFile, []byte(result.Text), 0644)
	if err != nil {
		log.Fatalf("Failed to save output: %v", err)
	}

	fmt.Printf("\n✅ Output saved to: %s\n", outputFile)
}
