package problem1122

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type input struct {
	arr1 []int
	arr2 []int
}

type testUnit struct {
	input  input
	output []int
}

func TestRelativeSortArray(t *testing.T) {
	table := []testUnit{
		{
			input{
				[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
				[]int{2, 1, 4, 3, 9, 6},
			},
			[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}

	for _, tt := range table {
		if actual := relativeSortArray(tt.input.arr1, tt.input.arr2); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("relativeSortArray(%v, %v) expect %v, but got %v", tt.input.arr1, tt.input.arr2, tt.output, actual)
		}
	}
}
