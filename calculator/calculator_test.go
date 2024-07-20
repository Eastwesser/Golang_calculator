package calculator

import (
	"testing"
)

// TestCalculateExpression тестирует функцию CalculateExpression для различных математических выражений
func TestCalculateExpression(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectPanic bool
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
		// Перехватываем панику
		defer func() {
			if r := recover(); r != nil {
				if !test.expectPanic {
					t.Errorf("CalculateExpression(%q) panicked with %v; did not expect panic", test.input, r)
				}
			}
		}()

		if test.expectPanic {
			func() {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("CalculateExpression(%q) did not panic; expected panic", test.input)
					}
				}()
				CalculateExpression(test.input)
			}()
		} else {
			result := CalculateExpression(test.input)
			if result != test.expected {
				t.Errorf("CalculateExpression(%q) = %s; want %s", test.input, result, test.expected)
			}
		}
	}
}
