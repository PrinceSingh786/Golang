package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}
type Byfield []person

func (b Byfield) Len() int {
	return len(b)
}
func (b Byfield) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b Byfield) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func main() {
	p1 := person{Name: "Prince", Age: 22}
	p2 := person{Name: "Jack", Age: 20}
	p3 := person{Name: "Almanac", Age: 18}
	people := []person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(Byfield(people))
	fmt.Println(people)

}
