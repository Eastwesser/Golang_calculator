package calculator

import (
	"testing"
)

func TestCalculateExpression(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		// Корректные римские случаи
		{"IX + II", "IX + II", "XI", false},
		{"IV - II", "IV - II", "II", false},
		{"X * II", "X * II", "X", false},
		{"X / II", "X / II", "V", false},
		{"C - X", "C - X", "XC", false},

		// Корректные арабские случаи
		{"10 + 5", "10 + 5", "15", false},
		{"10 - 5", "10 - 5", "5", false},
		{"10 * 5", "10 * 5", "50", false},
		{"10 / 5", "10 / 5", "2", false},
		{"5 + 3", "5 + 3", "8", false},
		{"8 - 5", "8 - 5", "3", false},
		{"7 * 2", "7 * 2", "14", false},
		{"6 / 3", "6 / 3", "2", false},

		// Ошибочные случаи
		{"10 / 0", "10 / 0", "", true},               // Деление на ноль
		{"10 + 11", "10 + 11", "", true},             // Неправильный диапазон для арабских чисел
		{"IX + 5", "IX + 5", "", true},               // Смешивание римских и арабских чисел
		{"IV - V", "IV - V", "", true},               // Результат отрицателен, а числа римские
		{"5 + II", "5 + II", "", true},               // Смешивание римских и арабских чисел
		{"IX - 3", "IX - 3", "", true},               // Смешивание римских и арабских чисел
		{"0 + 5", "0 + 5", "", true},                 // Неправильный диапазон для арабских чисел
		{"11 - 1", "11 - 1", "", true},               // Неправильный диапазон для арабских чисел
		{"X * XI", "X * XI", "", true},               // Неправильный диапазон для арабских чисел
		{"X / 3", "X / 3", "", true},                 // Неправильный формат входных данных (X не преобразуется в число)
		{"IV + IV + IV", "IV + IV + IV", "", true},   // Неправильный формат входных данных (три операнда)
		{"IV / II / III", "IV / II / III", "", true}, // Неправильный формат входных данных (больше двух операндов)
		{"5 *", "5 *", "", true},                     // Неправильный формат входных данных (отсутствует второй операнд)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateExpression(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("CalculateExpression() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if result != tt.expected {
				t.Errorf("CalculateExpression() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
