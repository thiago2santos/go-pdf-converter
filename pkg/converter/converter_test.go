package converter

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
		want   bool
	}{
		{
			name:   "with nil config",
			config: nil,
			want:   true,
		},
		{
			name:   "with custom config",
			config: &Config{OCRFallback: true, Verbose: true},
			want:   true,
		},
		{
			name:   "with default config",
			config: DefaultConfig(),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conv := New(tt.config)
			if (conv != nil) != tt.want {
				t.Errorf("New() returned nil, want non-nil")
			}
			if conv.config == nil {
				t.Errorf("New() config is nil, want non-nil")
			}
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config == nil {
		t.Fatal("DefaultConfig() returned nil")
	}

	if !config.OCRFallback {
		t.Error("DefaultConfig() OCRFallback = false, want true")
	}

	if config.Verbose {
		t.Error("DefaultConfig() Verbose = true, want false")
	}
}

func TestConfig(t *testing.T) {
	tests := []struct {
		name        string
		ocrFallback bool
		verbose     bool
	}{
		{
			name:        "OCR enabled and verbose",
			ocrFallback: true,
			verbose:     true,
		},
		{
			name:        "OCR disabled and quiet",
			ocrFallback: false,
			verbose:     false,
		},
		{
			name:        "OCR enabled and quiet",
			ocrFallback: true,
			verbose:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &Config{
				OCRFallback: tt.ocrFallback,
				Verbose:     tt.verbose,
			}

			if config.OCRFallback != tt.ocrFallback {
				t.Errorf("Config.OCRFallback = %v, want %v", config.OCRFallback, tt.ocrFallback)
			}

			if config.Verbose != tt.verbose {
				t.Errorf("Config.Verbose = %v, want %v", config.Verbose, tt.verbose)
			}
		})
	}
}

func TestResult(t *testing.T) {
	result := &Result{
		Text:            "Sample text",
		Method:          MethodTextExtraction,
		TotalPages:      3,
		LinesCount:      10,
		WordsCount:      50,
		CharactersCount: 300,
	}

	if result.Text != "Sample text" {
		t.Errorf("Result.Text = %v, want 'Sample text'", result.Text)
	}

	if result.Method != MethodTextExtraction {
		t.Errorf("Result.Method = %v, want %v", result.Method, MethodTextExtraction)
	}

	if result.TotalPages != 3 {
		t.Errorf("Result.TotalPages = %v, want 3", result.TotalPages)
	}

	if result.WordsCount != 50 {
		t.Errorf("Result.WordsCount = %v, want 50", result.WordsCount)
	}
}

func TestExtractionMethod(t *testing.T) {
	tests := []struct {
		name   string
		method ExtractionMethod
		want   string
	}{
		{
			name:   "text extraction method",
			method: MethodTextExtraction,
			want:   "Text Extraction",
		},
		{
			name:   "OCR method",
			method: MethodOCR,
			want:   "OCR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.method) != tt.want {
				t.Errorf("ExtractionMethod = %v, want %v", tt.method, tt.want)
			}
		})
	}
}

func TestConvertNonExistentFile(t *testing.T) {
	conv := New(nil)
	_, err := conv.Convert("non_existent_file.pdf")

	if err == nil {
		t.Error("Convert() with non-existent file should return error")
	}
}

func TestConvertInvalidFile(t *testing.T) {
	conv := New(nil)

	// Try to convert this Go source file as a PDF
	_, err := conv.Convert("converter_test.go")

	if err == nil {
		t.Error("Convert() with invalid PDF should return error")
	}
}

