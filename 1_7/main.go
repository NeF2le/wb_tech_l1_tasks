package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	m map[string]int
}

func main() {
	safeMap := &SafeMap{m: map[string]int{"a": 1}}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		safeMap.Lock()
		defer safeMap.Unlock()
		for range 10 {
			safeMap.m["a"]++
		}
	}()

	safeMap.Lock()
	safeMap.m["a"]++
	safeMap.Unlock()

	wg.Wait()
	fmt.Println(safeMap.m["a"])
}
