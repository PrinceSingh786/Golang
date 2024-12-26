package main

import "fmt"

type person struct {
	first              string
	last               string
	ice_cream_flavours []string
}

func main() {
	p1 := person{
		first:              "Prince",
		last:               "Singh",
		ice_cream_flavours: []string{"vanilla", "Choco", "Katori"},
	}
	fmt.Println(p1)
	for _, value := range p1.ice_cream_flavours {
		fmt.Println(value)
	}
}
