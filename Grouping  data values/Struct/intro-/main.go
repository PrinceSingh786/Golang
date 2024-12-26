//Struct allows us to compose together values of same data type

package main

import "fmt"

type emp struct {
	name string
	age  int
}
type manager struct {
	emp
	ismarried bool
}

func main() {
	e1 := struct {
		name string
		age  int
	}{
		name: "prince singh",
		age:  21,
	}
	e2 := emp{
		name: "Roderick",
		age:  29,
	}
	fmt.Printf("%T \t", e1)
	fmt.Printf("%T \t", e2)
	// m1 := manager{
	// 	emp:       e1,
	// 	ismarried: true,
	// }
	// e2 := emp{
	// 	name: "Alex Perry",
	// 	age:  32,
	// }
	// m2 := manager{
	// 	emp: emp{
	// 		name: "RAJA",
	// 		age:  28,
	// 	},
	// 	ismarried: true,
	// }
	// fmt.Println(m1)
	// fmt.Println(m2)
	// fmt.Println(m1.name)
	// fmt.Println(m1.emp.name + " is a hero.")
	// fmt.Printf("%T -- %#v\n", e1, e1)
	// fmt.Println(e1.name, e2.name)
}
