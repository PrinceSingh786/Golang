package main

import (
	"fmt"
)

func main() {

	vector := []int{1, 2, 3, 4, 55, 6, 66, 7, 777, 7777}
	fmt.Println(vector)

	vector = append(vector[:3], vector[5:]...)
	fmt.Println(vector)

}
