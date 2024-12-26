/*
JSON --javascript object notation
It is a lightweight data interchange format between a client and a server.

Key features ---
   1. Human-Readable( simple text format)
   2. Lightweight (compact and efficient format)
   3. Language-Independent (supported by almost all programming languages.)

JSON structure
    Objects : key-value pairs
	Arrays  : An ordered list of values in []
*/

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "Prince",
		Age:  22,
	}
	p2 := person{
		Name: "Jack",
		Age:  21,
	}
	people := []person{p1, p2}
	fmt.Println(people)
	encoded, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encoded))
}
