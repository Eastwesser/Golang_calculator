package roman

import (
	"testing"
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
		{"MMM", 3000, false},
		{"MMMCMXCIX", 3999, false},
		{"MMMM", 0, true},
		{"INVALID", 0, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := RomanToArabic(test.input)
			if test.isError {
				if err == nil {
					t.Errorf("Expected an error for input: %s, but got none", test.input)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error for input: %s, but got: %v", test.input, err)
				} else if result != test.expected {
					t.Errorf("Unexpected result for input: %s. Got: %d, Expected: %d", test.input, result, test.expected)
				}
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
			if result != test.expected {
				t.Errorf("Unexpected result for input: %d. Got: %s, Expected: %s", test.input, result, test.expected)
			}
		})
	}
}
