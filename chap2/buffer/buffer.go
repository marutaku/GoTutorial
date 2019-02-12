package main

import "fmt"
import "time"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	c <- 3
	fmt.Println(<-c)
	c <- 4
	c <- 5
	go func() {
		time.Sleep(3 * time.Second)
		<-c
	}()
	c <- 6
	fmt.Println(<-c)
	fmt.Println(<-c)
}
