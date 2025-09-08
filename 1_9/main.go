package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ch1 := make(chan int, len(nums))
	ch2 := make(chan int, len(nums))

	go func() {
		for _, n := range nums {
			ch1 <- n
		}
		close(ch1)
	}()

	go func() {
		for v := range ch1 {
			ch2 <- v * 2
		}
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
}
