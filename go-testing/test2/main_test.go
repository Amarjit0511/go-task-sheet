package main

import "testing"

func TestPrintln(t *testing.T) {
	result := print()
	expected := "Amarjt"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
