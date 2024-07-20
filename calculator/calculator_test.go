package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		{"10 + X", "", true},
		{"X + 10", "", true},
		{"5 / 0", "", true},
		{"", "", true},
		{"5 +", "", true},
		{"3 % 2", "", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := CalculateExpression(test.input)
			if test.isError {
				require.Error(t, err, "Expected an error for input: %s", test.input)
			} else {
				require.NoError(t, err, "Did not expect an error for input: %s", test.input)
				assert.Equal(t, test.expected, result, "Unexpected result for input: %s", test.input)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		x        int
		y        int
		operator string
		expected int
		isError  bool
	}{
		{3, 7, "+", 10, false},
		{10, 4, "-", 6, false},
		{2, 5, "*", 10, false},
		{8, 2, "/", 4, false},
		{8, 0, "/", 0, true},
		{5, 3, "%", 0, true},
	}

	for _, test := range tests {
		t.Run(test.operator, func(t *testing.T) {
			result, err := calculate(test.x, test.y, test.operator)
			if test.isError {
				require.Error(t, err, "Expected an error for operator: %s", test.operator)
			} else {
				require.NoError(t, err, "Did not expect an error for operator: %s", test.operator)
				assert.Equal(t, test.expected, result, "Unexpected result for operator: %s", test.operator)
			}
		})
	}
}
