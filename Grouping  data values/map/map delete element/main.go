package main

import "fmt"

func main() {
	mp := map[string]string{
		"prince": "singh",
		"Alex":   "Todd",
		"Mary":   "Jane",
	}

	delete(mp, "prijnce")
	fmt.Println(mp["prince"])
	fmt.Println(mp)
}
