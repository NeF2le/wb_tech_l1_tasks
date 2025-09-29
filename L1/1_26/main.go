package main

import (
	"fmt"
	"strings"
)

func isUniqueSymbols(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]struct{})
	for _, v := range str {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(isUniqueSymbols("abcde"))
	fmt.Println(isUniqueSymbols("Hello World"))
}
