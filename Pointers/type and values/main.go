// &  is for getting the address of the variable
// * is for deferencing the address and getting the value.
package main

import (
	"fmt"
)

func main() {
	x := 4
	y := &x
	fmt.Printf("%v has type %T \n", &x, &x)
	fmt.Printf("%v has type %T \n", *&*y, *y)
	*y = 444
	fmt.Println(x)

}
