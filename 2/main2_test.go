package main

import "testing"

func TestSliceExample(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sliceExample(input)

	expected := []int{2, 4, 6, 8, 10}

	if len(result) != len(expected) {
		t.Errorf("sliceExample() lenght = %d, want %d", len(result), len(expected))
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("sliceExample()[%d] = %d, want %d", i, result[i], expected[i])
		}
	}
}

func TestAddElements(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := 666

	result := addElements(input, num)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 666}

	if len(result) != len(expected) {
		t.Errorf("addElements() lenght = %d,want %d", len(result), len(expected))
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("addElements()[%d] = %d,want %d", i, result[i], expected[i])
		}
	}
}

func TestCopySlice(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := copySlice(input)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if len(input) != len(expected) {
		t.Errorf("copySlice() lenght = %d,want %d", result, expected)
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("copySlice()[%d] = %d,want %d", i, result[i], expected[i])
		}
	}

	input[0] = 999
	if result[0] == 999 {
		t.Errorf("copySlice() — копия зависит от оригинала")
	}

}

func TestRemoveElement(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	idx := 3
	result := removeElement(input, idx)

	expected := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}

	if len(result) != len(expected) {
		t.Errorf("removeElement() lenght = %d,want %d", len(result), len(expected))
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("copySlice()[%d] = %d,want %d", i, result[i], expected[i])
		}
	}

}
