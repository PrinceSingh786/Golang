package main

import (
	"fmt"
)

func main() {
	fmt.Println("2+3 =", mysum(2, 3))
}
func mysum(xi ...int) int {
	sum := 1
	for _, v := range xi {
		sum += v
	}
	return sum
}
