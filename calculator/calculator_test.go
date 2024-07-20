package calculator

import (
	"testing"
)

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
		{"X / 0", "", true},
		{"5 % 2", "", true},
		{"3 + 11", "", true},
		{"11 + 3", "", true},
		{"", "", true},
		{"3 +", "", true},
		{"3 3 +", "", true},
		{"X + X", "XX", false},
		{"X + 1", "", true},
		{"I + X", "", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := CalculateExpression(test.input)
			if test.isError {
				if err == nil {
					t.Errorf("Expected an error for input: %s, but got none", test.input)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error for input: %s, but got: %v", test.input, err)
				} else if result != test.expected {
					t.Errorf("Unexpected result for input: %s. Got: %s, Expected: %s", test.input, result, test.expected)
				}
			}
		})
	}
}
