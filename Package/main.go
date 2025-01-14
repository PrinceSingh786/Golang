package main

import (
	"fmt"
	"\\Golang\\Package\dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "fido",
		age:  dog.years,
	}
	fmt.Println(fido)
}