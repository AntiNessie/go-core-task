package main

import "fmt"

func main() {

	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(cross(a, b))

}

func cross(a, b []int) (bool, []int) {

	map1 := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range a {
		map1[v] = true
	}

	for _, v := range b {
		if map1[v] {
			result = append(result, v)
		}
	}

	return len(result) > 0, result

}
