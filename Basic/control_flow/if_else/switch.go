package main

import "fmt"

func main() {
	x := 42
	switch x {
	case 40:
		fmt.Println("mai 40 hu")
	case 41:
		fmt.Println("mai 41 hu")
	case 42:
		fmt.Println("mai 42 hu")
	case 43:
		fmt.Println("mai 43 hu")
	case 44:
		fmt.Println("mai 44 hu")
	default:
		fmt.Println("mai 40 aur 44 k range k bahar hu")
	}
}
