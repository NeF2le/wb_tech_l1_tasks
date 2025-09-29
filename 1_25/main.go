package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	fmt.Println("Program Start")
	sleep(time.Second * 5)
	fmt.Println("Program End")
}
