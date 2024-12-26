package main

import "fmt"

func main() {
	vector := make([]int, 0, 5)
	fmt.Println(vector)
	fmt.Println(len(vector))
	fmt.Println(cap(vector))
	vector = append(vector, 1, 2, 4, 5, 33, 34, 435, 34)
	fmt.Println("--------")
	fmt.Println(vector)
	fmt.Println(len(vector))
	fmt.Println(cap(vector))
}
