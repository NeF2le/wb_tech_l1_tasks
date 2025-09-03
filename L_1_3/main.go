package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	n := flag.Int("workersNum", 5, "workers number")
	flag.Parse()

	ch := make(chan int)

	for i := 1; i <= *n; i++ {
		go func() {
			for v := range ch {
				fmt.Printf("worker[%d]: %d\n", i, v)
			}
		}()
	}

	for {
		ch <- rand.Intn(100)
	}
}
