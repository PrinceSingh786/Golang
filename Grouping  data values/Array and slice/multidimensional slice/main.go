package main

import "fmt"

func main() {
	vector1 := []int{1, 2, 4, 5}
	vector2 := []int{10, 20, 40, 50}
	vector3 := []int{100, 200, 400, 500}
	d2 := [][]int{vector1, vector2, vector3}
	fmt.Println(d2)
}
