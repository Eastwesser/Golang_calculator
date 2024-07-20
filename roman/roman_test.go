package roman

import (
	"testing"
)

func TestRomanToNumeral(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"X", true},
		{"A", false},
		{"IV", true},
		{"", false},
		{"XIII", true},
		{"INVALID", false},
	}

	for _, test := range tests {
		result := RomanToNumeral(test.input)
		if result != test.expected {
			t.Errorf("Expected %v, got %v for input %s", test.expected, result, test.input)
		}
	}
}

func TestRomanToArabic(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"X", 10},
		{"IV", 4},
		{"IX", 9},
		{"I", 1},
		{"V", 5},
		{"XX", 20},
	}

	for _, test := range tests {
		result := RomanToArabic(test.input)
		if result != test.expected {
			t.Errorf("Expected %d, got %d for input %s", test.expected, result, test.input)
		}
	}
}

func TestArabicToRoman(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{10, "X"},
		{4, "IV"},
		{9, "IX"},
		{1, "I"},
		{5, "V"},
		{20, "XX"},
	}

	for _, test := range tests {
		result := ArabicToRoman(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s for input %d", test.expected, result, test.input)
		}
	}
}

func TestArabicToRomanComplex(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1987, "MCMLXXXVII"},
		{3999, "MMMCMXCIX"},
		{44, "XLIV"},
		{49, "XLIX"},
		{2022, "MMXXII"},
	}

	for _, test := range tests {
		result := ArabicToRoman(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s for input %d", test.expected, result, test.input)
		}
	}
}
