package main

import "fmt"

type employee struct {
	name      string
	age       int
	ismarried bool
}

func (e employee) intro() {
	fmt.Println("I am ", e.name)
	if e.ismarried {
		fmt.Println("I am married")
	} else {
		fmt.Println("I am not married")
	}
	fmt.Println("i am ", e.age)
}
func main() {
	e1 := employee{
		name:      "prince singh",
		age:       21,
		ismarried: false,
	}
	e2 := employee{
		name:      "Alex perry",
		age:       31,
		ismarried: true,
	}
	e1.intro()
	e2.intro()
}
