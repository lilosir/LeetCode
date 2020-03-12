package problem977

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type testUnit struct {
	input  []int
	output []int
}

func TestSortedSquares(t *testing.T) {
	table := []testUnit{
		{
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			[]int{-7, -3, 2, 3, 11},
			[]int{4, 9, 9, 49, 121},
		},
	}

	for _, tt := range table {
		if actual := sortedSquares(tt.input); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("sortedSquares(%v), expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
