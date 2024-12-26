package main

import (
	"fmt"
)

func main() {
	// vector := []int{1, 2, 3, 4, 5}
	// fmt.Println(sum(vector))

	x := sum([]int{2, 4, 6, 8, 10})
	fmt.Println(x)
}

func sum(ii []int) (total int) {
	total = 0
	for _, value := range ii {
		total += value
	}
	return total
}
