package main

import (
	"fmt"
)

func main() {
	// Once a variable declared , you have to use it ceratainly.
	var x int = 42
	var y int
	fmt.Println(x)
	fmt.Println(y)
	y = 4
	b, c := 4, 34
	fmt.Println(b, c)
	fmt.Println(y)
}

/*
Go is statically typed language and not dynamic
declaring variable of certain type eliminate errrors .
It's like validating data.
code is specific and precise.
*/
