package calculator

import (
	"testing"
)

// TestCalculateExpression тестирует функцию CalculateExpression для различных математических выражений
func TestCalculateExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isError  bool
	}{
		{"3 + 7", "10", false},
		{"10 - 4", "6", false},
		{"2 * 5", "10", false},
		{"8 / 2", "4", false},
		{"X + V", "XV", false},
		{"X - V", "V", false},
		{"X * II", "XX", false},
		{"X / II", "V", false},
		{"10 + X", "", true},
		{"X + 10", "", true},
		{"5 / 0", "", true},
	}

	for _, test := range tests {
		result, err := CalculateExpression(test.input)
		if test.isError {
			if err == nil {
				t.Errorf("CalculateExpression(%q) expected error; got result %s", test.input, result)
			}
		} else {
			if err != nil || result != test.expected {
				t.Errorf("CalculateExpression(%q) = %s, %v; want %s, nil", test.input, result, err, test.expected)
			}
		}
	}
}
