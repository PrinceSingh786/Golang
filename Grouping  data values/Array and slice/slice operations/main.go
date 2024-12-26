package main

import "fmt"

func main() {
	vector := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(vector); i++ {
		fmt.Println(vector[i])
	}

}
