package main

import "fmt"

func main() {
	array := []int{4, 2, 42, 11, 53}
	for ind, value := range array {
		fmt.Println("Ranging over array", ind, value)
	}

	//map ranges
	m := map[int]string{
		2: "Prince",
		3: "Singh",
		4: "MMMUT",
	}
	for first, second := range m {
		fmt.Println("Ranging over map", first, second)
	}

}
