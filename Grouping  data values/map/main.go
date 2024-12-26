package main

import "fmt"

func main() {
	mp := map[string]string{
		"prince": "singh",
		"Alex":   "Todd",
		"Mary":   "Jane",
	}
	mp["raj"] = "kuwar"
	fmt.Println(mp["prince"])
	fmt.Printf("%#v\n", mp)
	for key, value := range mp {
		fmt.Println(key, value)
	}
	for _, value := range mp {
		fmt.Println(value)
	}
	for key := range mp {
		fmt.Println(key)
	}
	// mpp := make(map[string]string)
	// mpp["Prince"] = "Singh"
	// mpp["alex"] = "todd"
	// mpp["mary"] = "jane"
	// fmt.Println(mpp)
	// fmt.Println(len(mpp))
}
