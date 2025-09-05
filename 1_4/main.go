package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	n := flag.Int("workersNum", 5, "workers number")
	flag.Parse()

	ch := make(chan int)
	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for i := 1; i <= *n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for v := range ch {
				fmt.Printf("worker[%d]: %d\n", id, v)
			}
		}(i)
	}

loop:
	for {
		select {
		case <-ctx.Done():
			close(ch)
			break loop
		case ch <- rand.Intn(100):
		}
	}

	wg.Wait()
	fmt.Println("done")
}
