package main

import "fmt"

func main() {
	vector1 := []int{1, 2, 3, 4}
	// vector2 := vector1
	// fmt.Println(vector1)
	// fmt.Println(vector2)
	// vector1[1] = 20
	// fmt.Println(vector1)
	// fmt.Println(vector2)
	// vector2[1] = 200
	// fmt.Println(vector1)
	// fmt.Println(vector2)

	vector3 := make([]int, 4)
	copy(vector3, vector1)
	vector3[1] = 0
	fmt.Println(vector1)
	fmt.Println(vector3)

}
