package main

func main() {
	// for i := 0; i < 11; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			println(i, " has inner loop as ", j)
		}
	}
}
