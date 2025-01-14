package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	x := Greet("Prince")
	if x != " Welcome Prince" {
		t.Error("Expected", "Welcome Prince", "Got", x)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Prince"))
	//Output:
	// Welcome Prince
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Prince")
	}
}
