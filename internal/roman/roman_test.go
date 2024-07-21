package roman

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Тестируем функцию RomanToArabic
func TestRomanToArabic(t *testing.T) {
	// Определяем набор тестовых случаев
	testCases := []struct {
		name  string // Название теста
		input string // Входные данные для теста
	}{
		{"Недопустимое римское число с неизвестными символами", "IIII"}, // Тест на недопустимое римское число
		{"Недопустимое римское число с повторяющимися символами", "VV"}, // Тест на недопустимое римское число с повторениями
		{"Римское число вне диапазона", "XI"},                           // Тест на выход римского числа за пределы допустимого диапазона
	}

	// Запускаем каждый тестовый случай
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Проверяем, что функция вызывает панику для недопустимого ввода
			require.Panics(t, func() { RomanToArabic(tc.input) }, "Ожидалась паника для ввода: %s", tc.input)
		})
	}
}

// Тестируем функцию ArabicToRoman
func TestArabicToRoman(t *testing.T) {
	// Определяем набор тестовых случаев
	testCases := []struct {
		name  string // Название теста
		input int    // Входные данные для теста
	}{
		{"Арабское число вне диапазона (0)", 0},     // Тест на число 0
		{"Арабское число вне диапазона (101)", 101}, // Тест на число 101
	}

	// Запускаем каждый тестовый случай
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Проверяем, что функция вызывает панику для недопустимого ввода
			require.Panics(t, func() { ArabicToRoman(tc.input) }, "Ожидалась паника для ввода: %d", tc.input)
		})
	}
}
