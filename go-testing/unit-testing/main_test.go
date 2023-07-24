package main

import "testing"

func TestAddition(t *testing.T) {
	result := addition(10, 5)
	expected := 15
	if result != expected {
		t.Errorf("Addition failed: expected %d, got %d", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	result := subtraction(10, 5)
	expected := 5
	if result != expected {
		t.Errorf("Subtraction failed: expected %d, got %d", expected, result)
	}
}

func TestDivision(t *testing.T) {
	result := division(10, 2)
	expected := 5
	if result != expected {
		t.Errorf("Division failed: expected %d, got %d", expected, result)
	}

	// Testing division by zero
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Error")
		}
	}()
	division(10, 0)
}

func TestMultiplication(t *testing.T) {
	result := multiplication(10, 5)
	expected := 50
	if result != expected {
		t.Errorf("Multiplication failed: expected %d, got %d", expected, result)
	}
}
