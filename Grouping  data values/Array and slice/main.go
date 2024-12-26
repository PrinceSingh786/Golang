package main

import (
	"fmt"
)

func main() {
	/*a := [3]int{1, 2, 3}
	fmt.Println(a) */
	b := []string{"Hello", "Prince", "Singh", "Gorakhpur"}
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(b[0])
	// Blank identifier to not use a returned value
	for _, value := range b {
		fmt.Printf("index   %v\n", value)
	}
	// var c [2]string
	// c[0] = "Prince"
	// c[1] = "Singh"
	// fmt.Println(c)
}
