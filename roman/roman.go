package roman

import (
	"strings"
)

// Словарь для хранения значений римских чисел до 3999 (или выше, если нужно)
var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
	"CD": 400, "D": 500, "CM": 900, "M": 1000,
}

// RomanToNumeral проверяет, является ли строка допустимым римским числом
func RomanToNumeral(num string) bool {
	_, exists := romanNumerals[num]
	return exists
}

// RomanToArabic преобразует римское число в его арабское эквивалент
func RomanToArabic(num string) (int, bool) {
	value, exists := romanNumerals[num]
	return value, exists
}

// ArabicToRoman преобразует арабское число в его римское эквивалент
func ArabicToRoman(num int) string {
	var result strings.Builder

	romanValues := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, rv := range romanValues {
		for num >= rv.Value {
			result.WriteString(rv.Symbol)
			num -= rv.Value
		}
	}

	return result.String()
}
