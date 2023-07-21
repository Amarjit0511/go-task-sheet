package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello World"
	got := HelloWorld()

	if expected != got {
		t.Errorf("Expected String(%s) is not same as"+" actual string (%s)", expected, got)
	}
}
