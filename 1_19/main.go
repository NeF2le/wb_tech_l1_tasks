package main

import (
	"fmt"
	"github.com/rivo/uniseg"
	"strings"
)

func main() {
	s := "главрыба 😳😒🫖🇺🇦👨‍👩‍👧‍👦"

	var clusters []string
	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		clusters = append(clusters, gr.Str())
	}

	for l, r := 0, len(clusters)-1; l < r; l, r = l+1, r-1 {
		clusters[l], clusters[r] = clusters[r], clusters[l]
	}
	fmt.Println(strings.Join(clusters, ""))
}
