package main

import "fmt"

func main() {
	vector := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(vector...))
}

func sum(ii ...int) int {
	s := 0
	for _, v := range ii {
		s += v
	}
	return s
}
