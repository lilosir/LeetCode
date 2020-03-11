package problem728

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type input struct {
	left  int
	right int
}

type testUnit struct {
	input  input
	output []int
}

func TestSelfDividingNumbers(t *testing.T) {
	table := []testUnit{
		{
			input{
				1,
				22,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
	}

	for _, tt := range table {
		if actual := selfDividingNumbers(tt.input.left, tt.input.right); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("selfDividingNumbers(%d,%d) expects %v, but got %v", tt.input.left, tt.input.right, tt.output, actual)
		}
	}
}
