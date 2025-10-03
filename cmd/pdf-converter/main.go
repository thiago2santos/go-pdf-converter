package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/thiago2santos/go-pdf-converter/pkg/converter"
)

const (
	version         = "1.0.0"
	outputFilePerms = 0o644
	minRequiredArgs = 2
)

const (
	exitCodeSuccess = 0
	exitCodeError   = 1
)

func main() {
	if len(os.Args) < minRequiredArgs {
		printUsage()
		os.Exit(exitCodeError)
	}

	// Handle flags
	if os.Args[1] == "--version" || os.Args[1] == "-v" {
		fmt.Printf("pdf-converter version %s\n", version)
		os.Exit(exitCodeSuccess)
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		os.Exit(exitCodeSuccess)
	}

	pdfPath := os.Args[1]

	if err := validatePDFPath(pdfPath); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Converting '%s'...\n", filepath.Base(pdfPath))

	// Create converter with verbose output
	config := converter.DefaultConfig()
	config.Verbose = true
	conv := converter.New(config)

	// Convert PDF
	result, err := conv.Convert(pdfPath)
	if err != nil {
		log.Fatalf("Error converting PDF: %v", err)
	}

	// Create output file
	outputPath := strings.TrimSuffix(pdfPath, filepath.Ext(pdfPath)) + "_extracted.txt"

	content := fmt.Sprintf("PDF: %s\nExtraction Method: %s\n\n%s",
		filepath.Base(pdfPath), result.Method, result.Text)

	err = os.WriteFile(outputPath, []byte(content), outputFilePerms)
	if err != nil {
		log.Fatalf("Error writing output: %v", err)
	}

	// Print results
	fmt.Printf("Conversion completed using %s\n", result.Method)
	fmt.Printf("Text saved to: %s\n", outputPath)
	fmt.Printf("Stats: %d pages, %d lines, %d words, %d characters\n",
		result.TotalPages, result.LinesCount, result.WordsCount, result.CharactersCount)
}

func printUsage() {
	fmt.Println("Usage: pdf-converter [options] <pdf_file>")
	fmt.Println("\nOptions:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println("  -v, --version  Show version information")
	fmt.Println("\nOutput will be saved as <pdf_file>_extracted.txt")
}

func printHelp() {
	fmt.Printf("pdf-converter v%s\n\n", version)
	fmt.Println("A fast PDF to text converter with OCR capabilities.")
	fmt.Println("\nUsage:")
	fmt.Println("  pdf-converter <pdf_file>")
	fmt.Println("\nOptions:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println("  -v, --version  Show version information")
	fmt.Println("\nExamples:")
	fmt.Println("  pdf-converter document.pdf")
	fmt.Println("  pdf-converter ~/Downloads/invoice.pdf")
	fmt.Println("\nFeatures:")
	fmt.Println("  • Automatic text extraction for text-based PDFs")
	fmt.Println("  • OCR fallback for image-based PDFs")
	fmt.Println("  • Clean, formatted output with page markers")
	fmt.Println("  • Statistics about extracted content")
	fmt.Println("\nFor more information, visit: https://github.com/thiago2santos/go-pdf-converter")
}

func validatePDFPath(path string) error {
	if path == "" {
		return fmt.Errorf("file path is required")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file '%s' not found", path)
	}

	if !strings.HasSuffix(strings.ToLower(path), ".pdf") {
		return fmt.Errorf("file '%s' is not a PDF file", path)
	}

	return nil
}
