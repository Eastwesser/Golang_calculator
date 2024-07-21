package roman

import (
	"testing"
)

func TestRomanToArabic(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		expected  int
		expectErr bool
	}{
		{"I", "I", 1, false},
		{"II", "II", 2, false},
		{"III", "III", 3, false},
		{"IV", "IV", 4, false},
		{"V", "V", 5, false},
		{"IX", "IX", 9, false},
		{"X", "X", 10, false},
		{"XI", "XI", 0, true},
		{"IIII", "IIII", 0, true},
		{"VV", "VV", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := RomanToArabic(tc.input)
			if (err != nil) != tc.expectErr {
				t.Errorf("RomanToArabic() error = %v, expectErr %v", err, tc.expectErr)
				return
			}
			if result != tc.expected {
				t.Errorf("RomanToArabic() = %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestArabicToRoman(t *testing.T) {
	testCases := []struct {
		name      string
		input     int
		expected  string
		expectErr bool
	}{
		{"1", 1, "I", false},
		{"2", 2, "II", false},
		{"3", 3, "III", false},
		{"4", 4, "IV", false},
		{"5", 5, "V", false},
		{"9", 9, "IX", false},
		{"10", 10, "X", false},
		{"40", 40, "XL", false},
		{"50", 50, "L", false},
		{"90", 90, "XC", false},
		{"100", 100, "C", false},
		{"0", 0, "", true},
		{"101", 101, "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ArabicToRoman(tc.input)
			if (err != nil) != tc.expectErr {
				t.Errorf("ArabicToRoman() error = %v, expectErr %v", err, tc.expectErr)
				return
			}
			if result != tc.expected {
				t.Errorf("ArabicToRoman() = %v, want %v", result, tc.expected)
			}
		})
	}
}
