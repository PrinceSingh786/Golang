package main

import (
	"fmt"
)

type cat struct {
	name string
}

func (c cat) sound() {
	fmt.Println("I am ", c.name, " meow meow !")
}

func (c *cat) run() {
	c.name = "Cutie"
	fmt.Println("I am ", c.name, " and I am running.")
}

type ii interface {
	sound()
	run()
}

func iiRun(x ii) {
	x.run()
}

func main() {
	c1 := cat{
		name: "Tukie",
	}
	c1.sound()
	c1.run()

	c2 := &cat{
		name: "Limi",
	}
	c2.sound()
	c2.run()
	iiRun(c2)

}
