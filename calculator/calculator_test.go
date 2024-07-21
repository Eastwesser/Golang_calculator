package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateExpression(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Invalid operation (divide by zero)", "10 / 0"},
		{"Invalid input format (out of range)", "11 + 1"},
		{"Mixed numeral systems", "IX + 5"},
		{"Negative result in Roman numeral system", "IV - V"},
		{"Unsupported operation (**) with Arabic numerals", "2 ** 2"},
		{"Incomplete expression", "2 +"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Panics(t, func() { CalculateExpression(tc.input) }, "Expected panic for input: %s", tc.input)
		})
	}
}
