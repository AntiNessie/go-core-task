package main

import "testing"

func TestCross(t *testing.T) {

	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, result := cross(a, b)
	expectedBool := true
	expectedSlice := []int{64, 3}

	if ok != expectedBool {
		t.Errorf("cross() bool - %v, want %v", ok, expectedBool)
	}

	if len(result) != len(expectedSlice) {
		t.Errorf("cross() lenght = %d, want = %d", len(result), len(expectedSlice))
		return
	}

	for i := range result {
		if result[i] != expectedSlice[i] {
			t.Errorf("cross()[%d] = %d, want %d", i, result[i], expectedSlice[i])
		}
	}
}

func TestCross_NotOk(t *testing.T) {

	a := []int{65, 3, 58, 678, 64}
	b := []int{63, 2, 4, 43}

	notOk, result := cross(a, b)
	expectedBool := false
	expectedSlice := []int{}

	if notOk != expectedBool {
		t.Errorf("cross() bool - %v, want %v", notOk, expectedBool)
	}

	if len(result) != len(expectedSlice) {
		t.Errorf("cross() lenght = %d, want = %d", len(result), len(expectedSlice))
		return
	}

	for i := range result {
		if result[i] != expectedSlice[i] {
			t.Errorf("cross()[%d] = %d, want %d", i, result[i], expectedSlice[i])
		}
	}

}
