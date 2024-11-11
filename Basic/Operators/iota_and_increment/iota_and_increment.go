package main

import "fmt"

const (
	c0 = iota
	c1
)
const (
	c2 = iota
	c3
	c4
)

func main() {
	fmt.Println(c0, c1, c2, c3, c4)
}
