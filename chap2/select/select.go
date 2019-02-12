package main

import "fmt"

// 無名関数はchannel`c`に値が入るまで待機
// 関数`fibonacci`は`x`に値が入る限り `case c <- x`以下を実行する
// 無名関数で`quit`に値が入った時に`fibonacci`の`<- quit`以下が実行される

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	fmt.Println("Stat fibonacci")
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
