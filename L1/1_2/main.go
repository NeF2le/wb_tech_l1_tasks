package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, v := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(math.Pow(float64(v), 2))
		}(v)
	}

	wg.Wait()
}
