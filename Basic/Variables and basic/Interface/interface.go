package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length  float64
	breadth float64
}
type circle struct {
	radius float64
}

func (r rectangle) Area() float64 {
	return r.length * r.breadth
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	Area() float64
}

func info(s shape) float64 {
	return s.Area()
}

func main() {
	c1 := circle{
		radius: 5,
	}
	r1 := rectangle{
		length:  10,
		breadth: 2,
	}
	fmt.Println(info(c1))
	fmt.Println(info(r1))
}
