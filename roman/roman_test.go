package roman

import (
	"testing"
)

// TestRomanToNumeral тестирует функцию RomanToNumeral, чтобы проверить, является ли строка допустимым римским числом
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
		{"XLII", true},
		{"MCMXCIV", true},
	}

	for _, test := range tests {
		result := RomanToNumeral(test.input)
		if result != test.expected {
			t.Errorf("RomanToNumeral(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

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
		{"INVALID", 0}, // Предполагается, что INVALID возвращает 0
	}

	for _, test := range tests {
		result, ok := RomanToArabic(test.input)
		if !ok || result != test.expected {
			t.Errorf("RomanToArabic(%q) = %d, %v; want %d, true", test.input, result, ok, test.expected)
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
