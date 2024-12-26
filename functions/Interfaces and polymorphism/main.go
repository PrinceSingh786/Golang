/* An interface in Go defines a set of method signatures.
Polymorphism is the ability of a Value of a certain type to also be of another type.package main
In Go values can be of more than one type
*/

package main

import "fmt"

type person struct {
	name string
}

func (p person) speak() {
	fmt.Println("I am ", p.name)
}

type human interface {
	speak()
}

func say(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Prince singh",
	}
	say(p1)
}
