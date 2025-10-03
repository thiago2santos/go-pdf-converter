package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
	}{
		{
			name:   "Success - with nil config uses defaults",
			config: nil,
		},
		{
			name:   "Success - with custom config",
			config: &Config{OCRFallback: true, Verbose: true},
		},
		{
			name:   "Success - with default config",
			config: DefaultConfig(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conv := New(tt.config)
			assert.NotNil(t, conv, "New() should return non-nil converter")
			assert.NotNil(t, conv.config, "converter config should be initialized")
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	require.NotNil(t, config, "DefaultConfig() should not return nil")
	assert.True(t, config.OCRFallback, "DefaultConfig() should enable OCR fallback")
	assert.False(t, config.Verbose, "DefaultConfig() should disable verbose by default")
}

func TestConvert_NonExistentFile(t *testing.T) {
	conv := New(nil)
	result, err := conv.Convert("non_existent_file.pdf")

	assert.Error(t, err, "Convert() should return error for non-existent file")
	assert.Nil(t, result, "Convert() should return nil result on error")
}

func TestConvert_InvalidFile(t *testing.T) {
	conv := New(nil)

	result, err := conv.Convert("converter_test.go")

	assert.Error(t, err, "Convert() should return error for invalid PDF")
	assert.Nil(t, result, "Convert() should return nil result on error")
}
