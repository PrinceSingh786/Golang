package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 332234483488
		close(c)
	}()
	value, ok := <-c
	fmt.Println(value, ok)

	// Channel is closed that's why c gives 0 now
	//fmt.Println(<-c)
}
