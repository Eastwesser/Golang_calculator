package roman

import (
	"strings"
)

// Словарь для хранения значений римских чисел
var romanNumerals = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
}

var romanValues = []struct {
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

// RomanToArabic преобразует римское число в его арабский эквивалент
func RomanToArabic(num string) (int, error) {
	num = strings.ToUpper(num) // Преобразуем строку в верхний регистр
	total := 0
	prevValue := 0
	repeatCount := 0
	lastChar := ""

	// Перебираем строку с конца, добавляем или вычитаем значение в зависимости от положения символа
	for i := len(num) - 1; i >= 0; i-- { // Здесь итерация идет по байтам (не рунам)
		char := string(num[i])
		value, exists := romanNumerals[char]
		if !exists {
			panic("недопустимое римское число")
			// Раскомментить return для тестов!
			//return 0, errors.New("недопустимое римское число")
		}
		// Проверка на повторение символов
		if char == lastChar {
			repeatCount++
			if repeatCount > 2 && (char == "I" || char == "X" || char == "C") {
				panic("недопустимое римское число")
				// Раскомментить return для тестов!
				//return 0, errors.New("недопустимое римское число")
			}
			if repeatCount > 0 && (char == "V" || char == "L" || char == "D") {
				panic("недопустимое римское число")
				// Раскомментить return для тестов!
				//return 0, errors.New("недопустимое римское число")
			}
		} else {
			repeatCount = 0
		}

		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
		lastChar = char
	}

	if total < 1 || total > 10 {
		panic("римское число должно быть от 1 до 10 включительно")
		// Раскомментить return для тестов!
		//return 0, errors.New("римское число должно быть от 1 до 10 включительно")
	}

	return total, nil
}

// ArabicToRoman преобразует арабское число в его римский эквивалент
func ArabicToRoman(num int) (string, error) {
	if num < 1 || num > 100 {
		panic("число должно быть в диапазоне от 1 до 100")
		// Раскомментить return для тестов!
		//return "", errors.New("число должно быть в диапазоне от 1 до 100")
	}
	var result strings.Builder

	/*
		Этот цикл итерирует по всем структурам в romanValues,
		которые содержат пары значение-символ
		(например, {100, "C"}, {90, "XC"}, и т.д.) в порядке убывания значений.
	*/
	for _, rv := range romanValues { // Первый цикл: перебираем каждую пару значение-символ римского числа
		/*
			Этот цикл добавляет символы к результату столько раз, сколько это возможно для текущего значения num.
			Например, если num равно 20, сначала добавится "X" (10), затем снова "X" (еще 10), и num станет 0.
		*/
		for num >= rv.Value { // Второй цикл: пока текущее число больше или равно значению rv.Value
			result.WriteString(rv.Symbol) // Добавляем римский символ к результату
			num -= rv.Value               // Вычитаем значение из текущего числа
		}
	}

	return result.String(), nil
}
