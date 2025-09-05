package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// #1: завершение через return
func returnInGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("returnInGoroutine running")
		time.Sleep(1 * time.Second)
		fmt.Println("returnInGoroutine done")
		return
	}()
}

// #2: отмена контекста
func contextDoneGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("contextDoneGoroutine done")
				return
			default:
				fmt.Println("contextDoneGoroutine running")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	time.Sleep(300 * time.Millisecond)
	cancel()
}

// #3: таймаут (тут через time.After, но можно через таймер или context WithTimeout)
func timeoutGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		timer := time.After(1 * time.Second)

		for {
			select {
			case <-timer:
				fmt.Println("timeoutGoroutine done")
				return
			default:
				fmt.Println("timeoutGoroutine running")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

// #4: runtime.Goexit()
func goExitGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goExitGoroutine running")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("goExitGoroutine done")
		runtime.Goexit()
	}()
}

// #5: обработка завершения программы
func shutdownGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("shutdownGoroutine done")
				return
			default:
				fmt.Println("shutdownGoroutine running")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(400 * time.Millisecond)
	stop()
}

// #6: отмена через канал-сигнал
func cancelChGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	done := make(chan struct{})

	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("cancelChGoroutine done")
				return
			default:
				fmt.Println("cancelChGoroutine running")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(300 * time.Millisecond)
	close(done)
}

// #7: атомик-флаг
func atomicFlagGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	var stopFlag int32

	go func() {
		defer wg.Done()
		for {
			if atomic.LoadInt32(&stopFlag) == 1 {
				fmt.Println("atomicFlagGoroutine done")
				return
			}
			fmt.Println("atomicFlagGoroutine running")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(300 * time.Millisecond)
	atomic.StoreInt32(&stopFlag, 1)
}

// #8: закрытие ресурса, который читает горутина (кроме канала может быть файл/соединение с БД/сервером)
func closeProcessChGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	ch := make(chan int)

	go func() {
		defer wg.Done()
		for range ch {
			fmt.Println("closeProcessChGoroutine running")
		}
		fmt.Println("closeProcessChGoroutine done")
	}()

	time.Sleep(300 * time.Millisecond)
	close(ch)
}

// #9: паника
func panicGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panicGoroutine done")
				return
			}
		}()
		fmt.Println("panicGoroutine running")
		time.Sleep(100 * time.Millisecond)
		panic("panicGoroutine panic")
	}()
}

// #10: прямое завершение программы
func forceExitGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("forceExitGoroutine running")
		time.Sleep(5 * time.Second)
		fmt.Println("forceExitGoroutine done")
		os.Exit(1)
	}()
}

// #11: обычное завершение
func defaultGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("defaultGoroutine running")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("defaultGoroutine done")
	}()
}

// #12: дедлок
func deadlockGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	ch := make(chan int)

	go func() {
		defer wg.Done()
		<-ch
	}()
}

func main() {
	wg := &sync.WaitGroup{}

	returnInGoroutine(wg)
	contextDoneGoroutine(wg)
	timeoutGoroutine(wg)
	goExitGoroutine(wg)
	shutdownGoroutine(wg)
	cancelChGoroutine(wg)
	atomicFlagGoroutine(wg)
	closeProcessChGoroutine(wg)
	panicGoroutine(wg)
	defaultGoroutine(wg)
	deadlockGoroutine(wg)
	forceExitGoroutine(wg)

	wg.Wait()
}
