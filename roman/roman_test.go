package roman

import (
	"testing"
)

func TestRomanToArabic(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    int
		expectError bool
	}{
		// Корректные случаи
		{"I", "I", 1, false},
		{"IV", "IV", 4, false},
		{"V", "V", 5, false},
		{"IX", "IX", 9, false},
		{"X", "X", 10, false},
		{"XL", "XL", 40, false},
		{"L", "L", 50, false},
		{"XC", "XC", 90, false},
		{"C", "C", 100, false},

		// Ошибочные случаи
		{"Invalid", "A", 0, true},              // Некорректное римское число
		{"Invalid combination", "IC", 0, true}, // Неправильная комбинация символов
		{"Empty string", "", 0, true},          // Пустая строка
		{"Invalid sequence", "XXXX", 0, true},  // Неправильная последовательность римских чисел
		{"Invalid format", "VV", 0, true},      // Повторяющиеся символы
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RomanToArabic(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("RomanToArabic() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if result != tt.expected {
				t.Errorf("RomanToArabic() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestArabicToRoman(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    string
		expectError bool
	}{
		// Корректные случаи
		{"1", 1, "I", false},
		{"4", 4, "IV", false},
		{"5", 5, "V", false},
		{"9", 9, "IX", false},
		{"10", 10, "X", false},
		{"40", 40, "XL", false},
		{"50", 50, "L", false},
		{"90", 90, "XC", false},
		{"100", 100, "C", false},

		// Ошибочные случаи
		{"Below range", 0, "", true},      // Неправильное значение (меньше 1)
		{"Above range", 101, "", true},    // Неправильное значение (больше 100)
		{"Negative number", -5, "", true}, // Неправильное значение (отрицательное число)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ArabicToRoman(tt.input)
			if (result == "" && !tt.expectError) || (result != "" && tt.expectError) {
				t.Errorf("ArabicToRoman() = %v, expectError %v", result, tt.expectError)
				return
			}
			if result != tt.expected {
				t.Errorf("ArabicToRoman() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
