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
	switch {
	case x == 41:
		fmt.Println("x is equal to 41")
	case x == 42:
		fmt.Println("x is equal to 42") // We don't have default fallthrough
		fallthrough
	case x == 43:
		fmt.Println("x is equal to 43")
		fallthrough
	case x == 44:
		fmt.Println("x is equal to 44")
		fallthrough
	case x == 45:
		fmt.Println("x is equal to 45")
	default:
		fmt.Println("x is not in the range of 41 and 44")

	}
}
