// The var variable declares a list of variables as in function argument lists, the type is last.
// Short declaration assignment doesn't work outside the function
// Outside the function every statement begins with a keyword like var,func and so on
package main

import "fmt"

var a, b, c bool
var aa int = 43

func main() {
	var x int
	const mm int = 297
	c = true
	fmt.Println(a, b, c, x, mm)
	fmt.Println(aa)
	fmt.Printf("%T \t %T \t %T ", a, b, mm)
}
