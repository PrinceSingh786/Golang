package main

import (
	"fmt"
)

func main() {
	// Method 1
	// for init ; condition; post {}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Method 2
	// for condition {}
	x := 20
	for x < 30 {
		println(x)
		x++
	}

	//Method 3
	// for {}     just like while in c/c++
	// break takes you out of the for loop
	for {
		println("I am x", x)
		if x < 0 {
			break
		}
		x--
	}
}
