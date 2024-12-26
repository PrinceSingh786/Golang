// Stringer interface
package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
	pages int
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title, " It has ", b.pages, " no. of pages.")
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is a number : ", strconv.Itoa(int(c)))
}
func main() {
	b := book{
		title: "King in the North.",
	}
	var n count = 42
	fmt.Println(n)
	fmt.Println(b)
}
