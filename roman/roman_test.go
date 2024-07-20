package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRomanToArabic(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		isError  bool
	}{
		{"I", 1, false},
		{"IV", 4, false},
		{"V", 5, false},
		{"IX", 9, false},
		{"X", 10, false},
		{"XL", 40, false},
		{"L", 50, false},
		{"XC", 90, false},
		{"C", 100, false},
		{"MMM", 0, true},     // Invalid
		{"INVALID", 0, true}, // Invalid
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := RomanToArabic(test.input)
			if test.isError {
				require.Error(t, err, "Expected an error for input: %s", test.input)
			} else {
				require.NoError(t, err, "Did not expect an error for input: %s", test.input)
				assert.Equal(t, test.expected, result, "Unexpected result for input: %s", test.input)
			}
		})
	}
}

func TestArabicToRoman(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{3999, "MMMCMXCIX"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			result := ArabicToRoman(test.input)
			assert.Equal(t, test.expected, result, "Unexpected result for input: %d", test.input)
		})
	}
}
