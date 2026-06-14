package main

import (
	"testing"
)

func TestToString(t *testing.T) {

	result := toString(42, 052, 0x2A)
	expected := "424242"

	if result != expected {
		t.Errorf("toString() = %q,want %q", result, expected)
	}
}

func TestToString_DiffTypes(t *testing.T) {

	result := toString(3.14, "Golang", true)
	expected := "3.14Golangtrue"

	if result != expected {
		t.Errorf("toString() = %q, want %q", result, expected)
	}

}

func TestToRune(t *testing.T) {

	result := toRune("Go")
	expected := []rune{'G', 'o'}

	if len(result) != len(expected) {
		t.Errorf("toRune() lenght = %d, want %d", len(result), len(expected))
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("toRune()[%d] = %q, want %q", i, result[i], expected[i])
		}
	}
}

func TestHashWithSalt(t *testing.T) {

	input := []rune("test")

	result1 := hashWithSalt(input)
	result2 := hashWithSalt(input)

	if len(result1) != 64 || len(result2) != 64 {
		t.Errorf("hashWithSalt() lenght = %d and %d, want 64", len(result1), len(result2))
	}

	if result1 != result2 {
		t.Errorf("hashWithSalt() must be the same")
	}

}
