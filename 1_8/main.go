package main

import (
	"flag"
	"fmt"
)

func main() {
	var n int64 = 5

	idx := flag.Uint("idx", 1, "bit index to change")
	value := flag.Uint("value", 1, "value to change")
	flag.Parse()

	if *value == 1 {
		n = n | (1 << *idx)
	} else if *value == 0 {
		n = n & ^(1 << *idx)
	}

	fmt.Println(n)
}
