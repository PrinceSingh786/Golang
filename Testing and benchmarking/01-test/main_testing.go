package main

import "testing"

func TestMysum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		{[]int{4, 2, 4}, 10},
		{[]int{43, 4}, 47},
		{[]int{84, 2}, 86},
		{[]int{28, 43}, 71},
	}
	for _, v := range tests {
		x := mysum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}
