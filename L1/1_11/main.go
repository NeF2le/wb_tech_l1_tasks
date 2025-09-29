package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{2, 3, 4, 9, 8, 7, 6, 5, 4}

	m := make(map[int]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}

	var result []int
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}

	fmt.Println(result)
}
