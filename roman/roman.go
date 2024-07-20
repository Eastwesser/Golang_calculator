package roman

import (
	"strings"
)

// Словарь для хранения значений римских чисел
var romanNumerals = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
	"CD": 400, "D": 500, "CM": 900, "M": 1000,
}

// RomanToNumeral проверяет, является ли строка допустимым римским числом
func RomanToNumeral(num string) bool {
	if num == "" {
		return false
	}

	for i := 0; i < len(num); {
		valid := false
		for length := 1; length <= 2; length++ {
			if i+length <= len(num) {
				if _, exists := romanNumerals[num[i:i+length]]; exists {
					valid = true
					i += length
					break
				}
			}
		}
		if !valid {
			return false
		}
	}
	return true
}

// RomanToArabic преобразует римское число в его арабский эквивалент
func RomanToArabic(num string) (int, bool) {
	if !RomanToNumeral(num) {
		return 0, true
	}

	total := 0
	prevValue := 0

	for i := len(num) - 1; i >= 0; i-- {
		value := romanNumerals[string(num[i])]
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total, true
}

// ArabicToRoman преобразует арабское число в его римский эквивалент
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
