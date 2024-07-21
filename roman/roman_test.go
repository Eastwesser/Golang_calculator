package roman

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToArabic(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Invalid Roman numeral with unknown characters", "IIII"},
		{"Invalid Roman numeral with repeated characters", "VV"},
		{"Roman numeral out of range", "XI"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Panics(t, func() { RomanToArabic(tc.input) }, "Expected panic for input: %s", tc.input)
		})
	}
}

func TestArabicToRoman(t *testing.T) {
	testCases := []struct {
		name  string
		input int
	}{
		{"Arabic number out of range (0)", 0},
		{"Arabic number out of range (101)", 101},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Panics(t, func() { ArabicToRoman(tc.input) }, "Expected panic for input: %d", tc.input)
		})
	}
}
