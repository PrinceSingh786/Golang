package main

import (
	"fmt"
)

var (
	a   int       = 3
	b   int       = 4
	cmp complex64 = 4 + 2i
)

func main() {
	c := float32(a*a + b*b)
	fmt.Println(a, b, cmp, c, float64(43))
}
