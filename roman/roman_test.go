package roman

import (
	"testing"
)

func TestIsRomanNumeral(t *testing.T) {
	if !IsRomanNumeral("X") {
		t.Errorf("Expected true, got false")
	}
	if IsRomanNumeral("A") {
		t.Errorf("Expected false, got true")
	}
}

func TestRomanToArabic(t *testing.T) {
	result := RomanToArabic("X")
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}

func TestArabicToRoman(t *testing.T) {
	result := ArabicToRoman(10)
	if result != "X" {
		t.Errorf("Expected X, got %s", result)
	}
}

func TestArabicToRomanComplex(t *testing.T) {
	result := ArabicToRoman(1987)
	expected := "MCMLXXXVII"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
