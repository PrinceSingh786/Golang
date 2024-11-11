package main

import "fmt"

func main() {
	a := "4334"
	// type is inferred from the value on the RHS
	var x = true
	fmt.Printf("Type of %v is %T\n", a, a)
	fmt.Printf("Type of %v is %T", x, x)

}
