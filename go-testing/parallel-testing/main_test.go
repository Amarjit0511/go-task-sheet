package main

import "testing"

func TestPrintln(t *testing.T) {
	t.Run("Testing 1 in parallel", func(t *testing.T) {
		t.Parallel()
		result := print()
		expected := "Amarjit"
		if result != expected {
			t.Errorf("Result incorrect")
		}
	})

	t.Run("Testing 2 in parallel", func(t *testing.T) {
		t.Parallel()
		result := print()
		expected := "Amarjit"
		if result != expected {
			t.Errorf("Result not correct")
		}
	})
}
