package main

import "fmt"

// let's create our own type
type Bytesize int64

const (
	_           = iota // ignoring first value by assigning to blank identifier
	KB Bytesize = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d \t \t %b \n", KB, KB)
	fmt.Printf("%d \t \t %b \n", MB, MB)
	fmt.Printf("%d \t \t %b \n", GB, GB)
	fmt.Printf("%d \t \t %b \n", TB, TB)
	fmt.Printf("%d \t \t %b \n", PB, PB)
	fmt.Printf("%d \t \t %b \n", EB, EB)

}
