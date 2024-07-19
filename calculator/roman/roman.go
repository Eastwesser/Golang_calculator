package roman

import (
	"errors"
	"strings"
)

var romanToNumber = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
	"CD": 400, "D": 500, "CM": 900, "M": 1000,
}

var numsToRoman = []struct {
	Value  int
	Symbol string
}{
	{Value: 1, Symbol: "I"},
	{Value: 4, Symbol: "IV"},
	{Value: 5, Symbol: "V"},
	{Value: 9, Symbol: "IX"},
	{Value: 10, Symbol: "X"},
	{Value: 40, Symbol: "XL"},
	{Value: 50, Symbol: "L"},
	{Value: 90, Symbol: "XC"},
	{Value: 100, Symbol: "C"},
	{Value: 400, Symbol: "CD"},
	{Value: 500, Symbol: "D"},
	{Value: 900, Symbol: "CM"},
	{Value: 1000, Symbol: "M"},
}

// Convert Roman numeral to Arabic number
func ToArabic(roman string) (int, error) {
	var total int
	i := 0
	for i < len(roman) {
		if i+1 < len(roman) && romanToNumber[roman[i:i+2]] > 0 {
			total += romanToNumber[roman[i:i+2]]
			i += 2
		} else if romanToNumber[roman[i:i+1]] > 0 {
			total += romanToNumber[roman[i:i+1]]
			i++
		} else {
			return 0, errors.New("invalid Roman numeral")
		}
	}
	return total, nil
}

// Convert Arabic number to Roman numeral
func ToRoman(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("number must be greater than 0")
	}
	var roman strings.Builder
	for _, entry := range numsToRoman {
		for num >= entry.Value {
			roman.WriteString(entry.Symbol)
			num -= entry.Value
		}
	}
	return roman.String(), nil
}
