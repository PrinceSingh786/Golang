package main

import (
	"errors"
	"fmt"
)

type sqrtErr struct {
	lat  string
	long string
	err  error
}

func (se sqrtErr) Error() string {
	return fmt.Sprintf("math error : %v %v %v ", se.lat, se.long, se.err)
}
func main() {
	_, err := sqrt(-43)
	if err != nil {
		log.Println(err)
	}
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		//e := errors.New("More coffee nahi chahiye")
		e := fmt.Errorf("No coffee needed", f)
		return 0, sqrtErr{"fjskfj", "fjflsj", e}
	}
	return 38, nil
}
