package main

import (
	"fmt"
)

type fake_int int

func main() {
	var x fake_int = 443
	fmt.Printf("%d \n", x)
	fmt.Printf("%T", x)
	var y int = 24
	y = int(x)
	fmt.Printf("%d \n", y)
	fmt.Printf("%T", y)
}
