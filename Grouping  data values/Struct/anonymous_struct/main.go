/*Anonymous types are indeterminate
They have not been declared as a type yet.
The compiler has flexibility with anonymous types.
You can assign as anonymous type to a variable declared as a certain type. If the assignment can occur, the compiler will figure it out;
the compiler will do an implicit conversion.
You cannot assign a named type to a different named type.
*/

package main

import "fmt"

type emp struct {
	name string
	age  int
}

func main() {
	e1 := struct {
		name string
		age  int
	}{
		name: "prince singh",
		age:  21,
	}
	e2 := emp{
		name: "Roderick",
		age:  29,
	}
	e1 = e2
	e2 = e1
	fmt.Println(e1)
	fmt.Printf("%T \t", e1)
	fmt.Printf("%T", e2)
}
