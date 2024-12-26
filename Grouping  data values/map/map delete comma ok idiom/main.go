package main

import "fmt"

func main() {
	mp := map[string]int{
		"Prince": 1,
		"Singh":  2,
		"Alex":   3,
		"Lucas":  4,
	}
	if value, okk := mp["ex"]; okk {
		fmt.Println(value)
	} else {
		fmt.Println("Keys doesn't exist")
	}
}
