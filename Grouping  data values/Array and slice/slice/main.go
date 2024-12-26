package main

import "fmt"

func main() {
	vector := []int{1, 3, 4, 43, 534, 532, 5322, 567, 5}

	fmt.Println(vector)

	// Add an element to slice
	/*vector = append(vector, 43, 888, 877, 9)
	fmt.Println(vector) */

	//Slicing a slice
	fmt.Printf("vector - %#v\n", vector)
	fmt.Printf("------------")

	// [inclusive : exclusive]
	fmt.Printf("vector - %#v\n", vector[0:4])
	fmt.Printf("------------")

}
