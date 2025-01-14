package main

import (
	"fmt"
)

func main() {
	// Can send data only and not receive data
	// c := make(chan<- int)
	//can receive only data
	c := make(<-chan int)

	fmt.Printf("%T\n", c)

}
