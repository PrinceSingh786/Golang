package main

import (
	f "first/function_returned_named.go"
	"fmt"
	"math"
	"math/rand"
)

func main() {
	f.Add(3, 43)
	fmt.Println("My random number is ", rand.Intn(5))
	fmt.Println("My random number is ", rand.Intn(5))
	fmt.Println("My random number is ", rand.Intn(5))
	fmt.Println("My random number is ", rand.Intn(5))
	fmt.Println(math.Sqrt(9))
	fmt.Println(math.Pi) //Pi is a constant , not a function
}
