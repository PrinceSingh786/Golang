package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//Send channel
	go foo(c)
	// Receive channel
	hoo(c)
	fmt.Println("ended successfully!")
}

func foo(c chan<- int) {
	c <- 48
}

func hoo(c <-chan int) {
	fmt.Println(<-c)
}
