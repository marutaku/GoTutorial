package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	fmt.Println("Sum array.  array: ", a[len(a)/2:])
	go sum(a[len(a)/2:], c)
	x := <-c
	fmt.Println("Result of summation:", x)

	fmt.Println("Sum array.  array: ", a[:len(a)/2])
	go sum(a[:len(a)/2], c)
	y := <-c
	fmt.Println("Result of summation:", y)

	fmt.Println(x, y, x+y)
}
