package main

import (
	"fmt"
	"sort"
)

func main() {
	vector := []int{1, 4, 2, 3, 2}
	words := []string{"boy", "apple", "dog", "cat"}
	sort.Ints(vector)
	fmt.Println(vector)
	sort.Strings(words)
	fmt.Println(words)
}
