package main

import "fmt"

func main() {
	mapped := map[int]string{
		0: "Prince",
		1: "Himanshu",
		2: "Gorakhpur",
	}
	if value, ok := mapped[2]; ok {
		fmt.Println("2 is mapped to", value)
	} else {
		fmt.Println("2 is not assigned", value)
	}
}
