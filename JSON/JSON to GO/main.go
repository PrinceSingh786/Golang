package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json : "Name"`
	Age  int    `json : "Age"`
}

func main() {
	s := `[{"Name":"Prince","Age":22},{"Name":"Jack","Age":21}]`
	bs := []byte(s)
	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", bs)
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
