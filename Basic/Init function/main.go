package main

import "fmt"

func init() {
	fmt.Println("I am INIT function")
}

func main() {
	x := 4
	fmt.Println("I am MAIN function", x)
}
