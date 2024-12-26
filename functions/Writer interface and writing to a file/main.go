package main

import (
	"bytes"
	"fmt"
)

// type person struct {
// 	name string
// }

// func (p person) writeOut(w io.Writer){
// 	, err := w.Write([]byte(p.name))
// 	return err
// }

// []byte is a slice of bytes , where each element represents a  single byte.
//  A byte buffer is a region of memory used to temporarily store a sequence of bytes.
// A byte buffer allows you to read and write bytes to and from the buffer.
// The purpose of a byte buffer is to provide a flexible and efficient way to work with sequence of bytes.
// It serves as a temporary storage for byte data and enables operations such  as reading , writing, appending, and resizing byte sequences.

func main() {
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// defer f.Close()
	// s := []byte("Hello Guyz")
	// _, er := f.Write(s)
	// if er != nil {
	// 	log.Fatalf("error %s ", er)
	// }

	b := bytes.NewBufferString("Hello guyz ")
	fmt.Println(b.String())
	b.WriteString("I am Coder")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("It's lusev day")
	fmt.Println(b.String())
	b.Write([]byte(" I am Prince."))
	fmt.Println(b.String())
}
