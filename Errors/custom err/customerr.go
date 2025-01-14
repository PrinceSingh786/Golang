package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v ", ce.info)
}

func main() {
	c1 := customErr{
		info: "I am a beautiful error",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Printf("foo ran - ", e, "\n I am a method in type error implemented by customErr")
}
