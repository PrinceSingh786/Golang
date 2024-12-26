package main

import "fmt"

type info struct {
	age       int
	ismarried bool
	dept      string
}

func main() {
	m1 := map[string]info{
		"Prince": {
			age:       21,
			ismarried: true,
			dept:      "Computer science",
		},
	}
	fmt.Println(m1["Prince"])
	fmt.Printf("%T ", m1["Prince"])

}
