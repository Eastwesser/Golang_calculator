package calculator

import (
	"testing"
)

func TestCalculateExpression(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{"IX + II", "IX + II", "XI", false},
		{"IV - II", "IV - II", "II", false},
		{"X * II", "X * II", "XX", false},
		{"X / II", "X / II", "V", false},
		{"II + III", "II + III", "V", false},
		{"II * V", "II * V", "X", false},
		{"10 + 5", "10 + 5", "15", false},
		{"10 - 5", "10 - 5", "5", false},
		{"10 * 5", "10 * 5", "50", false},
		{"10 / 5", "10 / 5", "2", false},
		{"10 / 0", "10 / 0", "", true},
		{"11 + 1", "11 + 1", "", true},
		{"IX + 5", "IX + 5", "", true},
		{"IV - V", "IV - V", "", true},
		{"IV * 2", "IV * 2", "", true},
		{"4 * II", "4 * II", "", true},
		{"2 ** 2", "2 ** 2", "", true},
		{"2 +", "2 +", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := CalculateExpression(tc.input)
			if (err != nil) != tc.expectErr {
				t.Errorf("CalculateExpression() error = %v, expectErr %v", err, tc.expectErr)
				return
			}
			if result != tc.expected {
				t.Errorf("CalculateExpression() = %v, want %v", result, tc.expected)
			}
		})
	}
}
