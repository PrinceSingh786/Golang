package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.name))
	return err
}

func main() {
	p := person{
		name: "Jenny",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	var b bytes.Buffer
	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}
