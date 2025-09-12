package main

import "fmt"

func reverse(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

func main() {
	s := "snow dog sun"
	b := []byte(s)
	reverse(b, 0, len(b)-1)

	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}

	fmt.Println(string(b))
}
