package myfunc

import (
	"testing"
)

func TestSub(t *testing.T) {
	input := [][]int{
		{-1, 8},
		{7, 1200},
		{5, -9},
		{-8, -12},
	}

	expect := []int{-9, -1193, 14, 4}

	for i := range expect {
		output := Sub(input[i][0], input[i][1])
		if output != expect[i] {
			t.Errorf("myfunc.Sub(%d, %d) = \"%d\", expect \"%d\"", input[i][0], input[i][1], output, expect[i])
		}
	}
}

func TestAdd(t *testing.T) {
	input := [][]int{
		{-1, 8},
		{7, 1200},
		{5, -9},
		{-8, -12},
	}

	expect := []int{7, 1207, -4, -20}

	for i := range expect {
		output := Add(input[i][0], input[i][1])
		if output != expect[i] {
			t.Errorf("myfunc.Add(%d, %d) = \"%d\", expect \"%d\"", input[i][0], input[i][1], output, expect[i])
		}
	}
}
