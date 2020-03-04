package problem1304

import (
	"testing"
  "github.com/lilosir/LeetCode/util"
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
	}

	for _, tt := range table {
		if actual := sumZero(tt.input); !util.compareTwoSlice(actual, tt.output) {
			t.Errorf("sumZero(%dd) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
