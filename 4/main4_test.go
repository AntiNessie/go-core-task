package main

import "testing"

func TestDifference(t *testing.T) {

	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := difference(slice1, slice2)
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	if len(result) != len(expected) {
		t.Errorf("difference() lenght = %d, want %d", len(result), len(expected))
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("difference()[%d] = %q, want %q", i, result[i], expected[i])
		}
	}
}

func TestDifference_NoDifference(t *testing.T) {
	slice1 := []string{"a", "b"}
	slice2 := []string{"a", "b", "c"}

	result := difference(slice1, slice2)

	if len(result) != 0 {
		t.Errorf("difference() length = %d, want 0", len(result))
	}
}
