package main

import (
	"fmt"
	"math/rand"
)

func main() {

	count := 1
	for i := 1; i <= 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iteration %v and count %v\n", i, count)
			count++
		}
	}
}
