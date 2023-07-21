package main

import (
	"testing"
)

func TestBio(t *testing.T) {
	name := "Amarjit"
	age := 23
	gender := "Male"
	expected := []string{"Amarjit", "23", "Male"}
	result := Bio(name, age, gender)
	if len(result) != len(expected) {
		t.Errorf("Unexpected length of result. Expected %d, but got %d", len(expected), len(result))
	}
}

func TestBioTwo(t *testing.T) {
	name := "Amarjit"
	age := 23
	gender := "Male"
	expected := []string{"Amarjit", "23", "Male"}
	result := Bio(name, age, gender)

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Unexpected value at index %d. Expected %s, but got %s", i, expected[i], v)
		}
	}
}
