package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS \t ", runtime.GOOS)
	fmt.Println("CPUs \t", runtime.NumCPU())
	fmt.Println("ARCH \t", runtime.GOARCH)
	fmt.Println("Goroutines \t", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go fo()
	wg.Wait()
	fmt.Println("NO of Goroutines ", runtime.NumGoroutine())
}
func foo() {
	fmt.Println("I am king")
	wg.Done()
}
func fo() {
	fmt.Println("I am ")
	wg.Done()
}
