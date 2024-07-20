package calculator

import (
	"testing"
)

// TestAdd тестирует функцию add с различными входными данными
func TestAdd(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{1, 2, 3},
		{10, 10, 20},
		{5, 5, 10},
		{0, 0, 0},
		{-1, -2, -3},
		{100, -50, 50},
	}

	for _, test := range tests {
		result := add(test.x, test.y)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.x, test.y, result, test.expected)
		}
	}
}

// TestSubtract тестирует функцию subtract с различными входными данными
func TestSubtract(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{2, 1, 1},
		{10, 5, 5},
		{5, 5, 0},
		{0, 0, 0},
		{0, 1, -1},
		{-1, -1, 0},
		{100, 50, 50},
	}

	for _, test := range tests {
		result := subtract(test.x, test.y)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d; want %d", test.x, test.y, result, test.expected)
		}
	}
}

// TestMultiply тестирует функцию multiply с различными входными данными
func TestMultiply(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{2, 3, 6},
		{10, 10, 100},
		{5, 0, 0},
		{0, 5, 0},
		{-1, -1, 1},
		{2, -3, -6},
		{-2, 3, -6},
	}

	for _, test := range tests {
		result := multiply(test.x, test.y)
		if result != test.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", test.x, test.y, result, test.expected)
		}
	}
}

// TestDivide тестирует функцию divide с различными входными данными
func TestDivide(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{6, 2, 3},
		{10, 5, 2},
		{5, 2, 2},
		{0, 1, 0},
		{-10, -2, 5},
		{9, 3, 3},
	}

	for _, test := range tests {
		if test.y == 0 {
			continue // Пропускаем тесты, где делитель равен нулю
		}
		result := divide(test.x, test.y)
		if result != test.expected {
			t.Errorf("Divide(%d, %d) = %d; want %d", test.x, test.y, result, test.expected)
		}
	}
}

// TestDivideByZero проверяет, что деление на ноль вызывает панику
func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Ожидалась panic при делении на ноль")
		}
	}()
	divide(1, 0)
}
