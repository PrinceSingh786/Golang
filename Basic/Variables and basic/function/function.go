package main

import "fmt"

func add(a int, b int) (int, int, string) {
	return a + b, a * a, "Prince"
}
func main() {
	fmt.Println(add(4, 244))
	// Nested function first add is going to return and then value is printed
}
