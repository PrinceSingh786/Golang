package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Anonymous function is running")
	}()

	func(age int) {
		fmt.Println("My age is ", age)
	}(21)
}
