// Stringer interface
package main

import (
	"fmt"
	"log"
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

func logInfo(s fmt.Stringer) {
	log.Println("Log from 138 ", s.String())
}

func main() {
	b := book{
		title: "King in the North.",
	}
	var n count = 42
	logInfo(n)
	logInfo(b)
}
