package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Value int
	mu    sync.Mutex
}

func (c *Counter) inc() {
	c.mu.Lock()
	c.Value++
	c.mu.Unlock()
}

func main() {
	c := Counter{}

	workers := 10
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			c.inc()
		}()
	}
	wg.Wait()

	fmt.Println(c.Value)
}
