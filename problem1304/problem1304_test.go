package problem1304

import (
	"testing"
)

type testUnit struct {
	input  int
	output []int
}

func TestSumZero(t *testing.T) {
	table := []testUnit{
		{
			5,
			[]int{-2, -1, 0, 1, 2},
		},
		{
			3,
			[]int{-1, 0, 1},
		},
		{
			1,
			[]int{0},
		},
		{
			2,
			[]int{-1, 1},
		},
		{
			4,
			[]int{-2, -1, 1, 2},
		},
	}

	for _, tt := range table {
		if actual := sumZero(tt.input); !compareTwoSlice(actual, tt.output) {
			t.Errorf("sumZero(%dd) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}

func compareTwoSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
