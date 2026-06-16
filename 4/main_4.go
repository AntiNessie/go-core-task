package main

import "fmt"

func main() {

	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(difference(slice1, slice2))

}

func difference(slice1, slice2 []string) []string {
	result := make([]string, 0)

	map1 := make(map[string]bool)
	for _, v := range slice2 {
		map1[v] = true
	}

	for _, v := range slice1 {
		if !map1[v] {
			result = append(result, v)
		}
	}

	return result
}
