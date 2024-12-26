package main

import "fmt"

type engine struct {
	electric bool
}
type vehicle struct {
	engine
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		model: "X1 500",
		doors: 4,
		color: "Blue",
	}
	fmt.Println(v1)
}
