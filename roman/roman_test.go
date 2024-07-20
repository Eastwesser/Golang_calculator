package roman

import (
	"testing"
)

// TestRomanToArabic тестирует функцию RomanToArabic, чтобы преобразовать римские числа в арабские числа
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
		{"XLII", 42},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		result, err := RomanToArabic(test.input)
		if err != nil || result != test.expected {
			t.Errorf("RomanToArabic(%q) = %d, %v; want %d, nil", test.input, result, err, test.expected)
		}
	}
}

// TestArabicToRoman тестирует функцию ArabicToRoman, чтобы преобразовать арабские числа в римские числа
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
		{42, "XLII"},
		{1994, "MCMXCIV"},
	}

	for _, test := range tests {
		result := ArabicToRoman(test.input)
		if result != test.expected {
			t.Errorf("ArabicToRoman(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}

// TestArabicToRomanComplex тестирует функцию ArabicToRoman с более сложными числами
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
			t.Errorf("ArabicToRoman(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
