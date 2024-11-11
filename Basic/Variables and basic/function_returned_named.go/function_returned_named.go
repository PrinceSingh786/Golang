package f

import (
	"fmt"
	"github.com/PrinceSingh786/animal"
)

// the return values can be named and then they will act like defined at the top of the function
func Add(a int, b int) (result int) {
	result = a + b
	return result
}
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(5, 243))
	fmt.Println(swap("prince", "singh"))
}
