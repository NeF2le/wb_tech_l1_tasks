package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	timer := time.After(5 * time.Second)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	go func() {
		defer close(ch)
		for {
			select {
			case <-timer:
				return
			case ch <- rand.Intn(100):
			}
		}
	}()

	wg.Wait()
}
