package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 4039
	}()

	c2 := make(chan string)
	go func() {
		c2 <- "Prince singh"
	}()
	fmt.Println(<-c)
	fmt.Println(<-c2)
}
