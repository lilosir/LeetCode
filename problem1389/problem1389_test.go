package problem1389

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type input struct {
	nums  []int
	index []int
}

type testUnit struct {
	input  input
	output []int
}

func TestCreateTargetArray(t *testing.T) {
	table := []testUnit{
		{
			input{
				[]int{0, 1, 2, 3, 4},
				[]int{0, 1, 2, 2, 1},
			},
			[]int{0, 4, 1, 3, 2},
		},
		{
			input{
				[]int{1, 2, 3, 4, 0},
				[]int{0, 1, 2, 3, 0},
			},
			[]int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range table {
		if actual := createTargetArray(tt.input.nums, tt.input.index); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("createTargetArray(%v, %v) expect %v, but got %v", tt.input.nums, tt.input.index, tt.output, actual)
		}
	}
}
