package main

import (
	"fmt"
)

func main() {
	// a := 5
	// fmt.Println(a)
	// change(&a)
	// fmt.Println(a)

	ii := []int{2, 4, 5, 6}
	fmt.Println(ii)
	f(ii)
	fmt.Println(ii)

}

func change(n *int) {
	*n = 44
}

func f(a []int) {
	a[0] = 555
}
