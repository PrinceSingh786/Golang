package main

import (
	"fmt"
)

func main() {
	e := make(chan int)
	o := make(chan int)
	q := make(chan int)
	// Send channel
	go send(e, o, q)
	// Receive channel
	receive(e, o, q)
	fmt.Println("Program executed successfully!")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
	close(e)
	close(o)
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel ", v)
		case v := <-o:
			fmt.Println("From the odd channel ", v)
		case v := <-q:
			fmt.Println("From the quit channel ", v)
			return
		}
	}
}
