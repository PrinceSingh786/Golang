package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	xy := bar()
	fmt.Println(xy())
	fmt.Printf("%T --xy -- \n", xy)
	fmt.Printf("%T --bar-- \n", bar)
	fmt.Printf("%T --foo-- \n", foo)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 108
	}
}
