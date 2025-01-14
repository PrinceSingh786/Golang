// Benchmarking allows us to measure the performance of the code
package saying

import "fmt"

func Greet(s string) string {
	return fmt.Sprint(" Welcome ", s)
}
