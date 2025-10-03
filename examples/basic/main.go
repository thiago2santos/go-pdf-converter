package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thiago2santos/go-pdf-converter/pkg/converter"
)

const (
	requiredArgs  = 2
	exitCodeError = 1
)

func main() {
	if len(os.Args) != requiredArgs {
		fmt.Println("Usage: go run main.go <pdf_file>")
		os.Exit(exitCodeError)
	}

	pdfPath := os.Args[1]

	// Create converter with default config
	conv := converter.New(nil)

	// Convert PDF
	result, err := conv.Convert(pdfPath)
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	}

	// Print results
	fmt.Printf("Conversion successful!\n")
	fmt.Printf("Method: %s\n", result.Method)
	fmt.Printf("Pages: %d\n", result.TotalPages)
	fmt.Printf("Words: %d\n", result.WordsCount)
	fmt.Printf("Characters: %d\n", result.CharactersCount)
	fmt.Printf("\nExtracted text:\n%s\n", result.Text)
}
