// An expression is a combination of values,variables, operators, and function calls that are evaluated to produce a single value.
// It can be as simple as a literal value or a variable, or it can involve complex operations and function invocations.

package main

import (
	"fmt"
)

func main() {
	x := func() {
		fmt.Println("Anonymous function is running")
	}

	x()

	y := func(age int) {
		fmt.Println("My age is ", age)
	}
	y(21)
}
