// Callback means passing a function as an argument in a function which is going to be used somewhere in the function.
// Call - action of invoking or executing a function.
// back - callback function is called back or invoked by the original function after it has completed its execution or reached a specific point in its code.

package main

import "fmt"

func main() {
	x := domath(3, 2, add)
	fmt.Println(x)
	y := domath(3, 2, sub)
	fmt.Println(y)

	fmt.Printf("add -- %T\n", add)
	fmt.Printf("sub -- %T\n", sub)
	fmt.Printf("domath -- %T\n", domath)
}

func domath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
