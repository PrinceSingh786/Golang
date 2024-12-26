// Persistently and patiently you are bound to succeed
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Program started ")
	wg.Add(2)
	go func() {
		fmt.Println("Mai chala re chala")
		wg.Done()
	}()
	go func() {
		fmt.Println("It's a beautiful day!")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Program ended ")
}
