package problem1051

import (
	"testing"
)

type testUnit struct {
	input  []int
	output int
}

func TestHeightChecker(t *testing.T) {
	table := []testUnit{
		{
			[]int{1, 1, 4, 2, 1, 3},
			3,
		},
		{
			[]int{5, 1, 2, 3, 4},
			5,
		},
		{
			[]int{1, 2, 3, 4, 5},
			0,
		},
	}
	for _, tt := range table {
		if actual := heightChecker(tt.input); actual != tt.output {
			t.Errorf("heightChecker(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
