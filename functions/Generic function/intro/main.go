package main

import (
	"fmt"
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float32) float32 {
	return a + b
}

type myrange interface {
	~int | float32 | float64
}

func addT[T myrange](a, b T) T {
	return a + b
}

type mine int

func main() {
	fmt.Println(addI(4, 9))
	fmt.Println(addF(4.22, 5.55))
	fmt.Println(addT(4, 9))
	fmt.Println(addT(4.22, 5.55))
	var m mine = 44
	fmt.Println(addT(m, 44))

}
