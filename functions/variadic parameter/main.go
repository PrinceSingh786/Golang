// Variadic parameter is a func which takes an unlimited number of arguments.
// The final incoming parameter in a function signature may have a type prefixed with ... A function with  such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.

package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
func sum(ii ...int) int {
	s := 0
	for _, value := range ii {
		s = s + value
	}
	return s
}
