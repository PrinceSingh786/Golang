package main

import (
	"fmt"
)

func mul(x int) int {
	return x * 5
}
func main() {
	ch := make(chan int)
	go func() {
		ch <- mul(3)
	}()
	fmt.Println(<-ch)
}
