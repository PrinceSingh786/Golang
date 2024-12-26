package main

import "fmt"

func main() {
	defer meow()
	bark()
}

func bark() {
	fmt.Println("Bho bho bho!")
}
func meow() {
	fmt.Println("Meow meow meow!")
	bark()
}
