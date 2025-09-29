package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	var result []string
	for k := range m {
		result = append(result, k)
	}
	fmt.Println(result)
}
