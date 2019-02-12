package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Printf(".")
			time.Sleep(100 * time.Millisecond)
		}
	}()
	queue := make(chan int, 3)

	time.Sleep(1 * time.Second)
	queue <- 1
	fmt.Println(1)

	time.Sleep(1 * time.Second)
	queue <- 2
	fmt.Println(2)

	time.Sleep(1 * time.Second)
	queue <- 3
	fmt.Println(3)

	go func() {
		time.Sleep(3 * time.Second)
		<-queue
	}()

	time.Sleep(1 * time.Second)
	// Lock
	queue <- 4
	fmt.Println(4)
}
