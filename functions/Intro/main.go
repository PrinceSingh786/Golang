// Everything in Go is Pass by value
// func(receiver)identifier(parameters)(returns){
//}

package main

import "fmt"

func main() {
	laugh("Prince ")
	s, i := add(4, 9)
	fmt.Println(s, i)
}
func laugh(s string) {
	fmt.Println(s + "is  laughing")
}

func add(a int, b int) (string, int) {
	return "sum ho gaya", a + b
}
