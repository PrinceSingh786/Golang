package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("This is an error prototype")
	fmt.Println(e)
	fmt.Printf("%T", e)
}
