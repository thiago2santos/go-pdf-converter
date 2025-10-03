package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/thiago2santos/go-pdf-converter/pkg/converter"
)

const (
	outputFilePerms = 0o644
	requiredArgs    = 2
	exitCodeError   = 1
)

func main() {
	if len(os.Args) != requiredArgs {
		fmt.Println("Usage: go run main.go <pdf_file>")
		os.Exit(exitCodeError)
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
	fmt.Printf("\n%s\n", strings.Repeat("=", 50))
	fmt.Printf("CONVERSION SUMMARY\n")
	fmt.Printf("%s\n", strings.Repeat("=", 50))
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
	err = os.WriteFile(outputFile, []byte(result.Text), outputFilePerms)
	if err != nil {
		log.Fatalf("Failed to save output: %v", err)
	}

	fmt.Printf("\nOutput saved to: %s\n", outputFile)
}
