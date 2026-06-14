package main

import (
	"fmt"
	"math/rand"
)

func main() {

	originalSlice := make([]int, 10)

	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	slice1 := sliceExample(originalSlice)
	fmt.Println(slice1)

	slice2 := addElements(originalSlice, 666)
	fmt.Println(slice2)

	slice3 := copySlice(originalSlice)
	fmt.Println(slice3)

	slice4 := removeElement(originalSlice, 3)
	fmt.Println(slice4)

}

func sliceExample(slice []int) []int {

	result := make([]int, 0)

	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result

}

func addElements(slice []int, num int) []int {

	return append(slice, num)
}

func copySlice(slice []int) []int {

	result := make([]int, len(slice))
	copy(result, slice)
	return result

}

func removeElement(slice []int, idx int) []int {

	if idx < 0 || idx >= len(slice) {
		return slice
	}

	return append(slice[:idx], slice[idx+1:]...)
}
