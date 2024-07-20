package roman

import (
	"errors"
	"strings"
)

var romanValues := []struct {
	Value  int
	Symbol string
}{
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

// Словарь для хранения значений римских чисел
var romanNumerals = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
}

// RomanToArabic преобразует римское число в его арабский эквивалент
func RomanToArabic(num string) (int, error) {
	num = strings.ToUpper(num) // Преобразуем строку в верхний регистр
	total := 0
	prevValue := 0

	// Перебираем строку с конца, добавляем или вычитаем значение в зависимости от положения символа
	for i := len(num) - 1; i >= 0; i-- {
		value, exists := romanNumerals[string(num[i])]
		if !exists {
			return 0, errors.New("недопустимое римское число")
		}
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total, nil
}

// ArabicToRoman преобразует арабское число в его римский эквивалент
func ArabicToRoman(num int) string {
	var result strings.Builder

	// Перебираем римские значения и добавляем их к результату
	for _, rv := range romanValues {
		for num >= rv.Value {
			result.WriteString(rv.Symbol)
			num -= rv.Value
		}
	}

	return result.String()
}
