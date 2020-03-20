package problem1356

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type testUnit struct {
	input  []int
	output []int
}

func TestSortByBits(t *testing.T) {
	table := []testUnit{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			[]int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
		{
			[]int{10000, 10000},
			[]int{10000, 10000},
		},
		{
			[]int{2, 3, 5, 7, 11, 13, 17, 19},
			[]int{2, 3, 5, 17, 7, 11, 13, 19},
		},
		{
			[]int{10, 100, 1000, 10000},
			[]int{10, 100, 10000, 1000},
		},
	}
	for _, tt := range table {
		if actual := sortByBits(tt.input); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("sortByBits(%v) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
