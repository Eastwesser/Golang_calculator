package roman

import (
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
}

func IsRomanNumeral(num string) bool {
	_, exists := romanNumerals[num]
	return exists
}

func RomanToArabic(num string) int {
	return romanNumerals[num]
}

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
