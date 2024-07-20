package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(2, 1)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	result := divide(6, 2)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
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
