// A wrapper function or Wrapper is a function that provides an additional layer of abstraction or functionality around an existing function or method.
//Logging - for adding logging statements before and after invoking the wrapped function.
//Timing and profiling - wrappers can be used to measure the execution time of functions.
// Authentication and authorization - Wrappers can handle authentication and authorization checks before executing the wrapped function .
//Error handling  - Wrappers can intercept errors returned by the wrapped function and transform them into a different error type or add more contextual information.

package main

import (
	"fmt"
	"log"
	"os"
)

// func Timedfn(f func()) {
// 	start := time.Now()
// 	f()
// 	elapsed := time.Since(start)
// 	fmt.Println("Elapsed time :", elapsed)
// }

// func f() {
// 	fmt.Println("f function started executing .")
// 	fmt.Println(" f function is being executed .")
// }

func main() {
	// Timedfn(f)
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("readFile in main: %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error in readfile func: %s", err)
	}
	return xb, nil
}
