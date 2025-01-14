package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		// Panic : send on closed channel
		// c <- 32
	}()
	for v := range c {
		fmt.Println(v)
	}

}
