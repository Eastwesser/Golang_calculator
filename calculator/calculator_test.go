package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{1, 2, 3},
		{10, 10, 20},
		{5, 5, 10},
		{0, 0, 0},
		{-1, -2, -3},
	}

	for _, test := range tests {
		result := add(test.x, test.y)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{2, 1, 1},
		{10, 5, 5},
		{5, 5, 0},
		{0, 0, 0},
		{0, 1, -1},
	}

	for _, test := range tests {
		result := subtract(test.x, test.y)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{2, 3, 6},
		{10, 10, 100},
		{5, 0, 0},
		{0, 5, 0},
		{-1, -1, 1},
	}

	for _, test := range tests {
		result := multiply(test.x, test.y)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{6, 2, 3},
		{10, 5, 2},
		{5, 2, 2},
		{0, 1, 0},
		{-10, -2, 5},
	}

	for _, test := range tests {
		result := divide(test.x, test.y)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for division by zero")
		}
	}()
	divide(1, 0)
}
