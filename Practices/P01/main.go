// Persistently and patiently you are bound to succeed
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Println("x and y are ", x, " and ", y)
	switch {
	case x < 5 && y < 5:
		fmt.Println("both are less than 5")
	case x == 5 && y == 5:
		fmt.Println("both are 5")
	case x > 5 && y > 5:
		fmt.Println("both are greater than 5")
	default:
		fmt.Println("one is less than or equal to 5 and other is greater than or equal to 5")
	}

}
