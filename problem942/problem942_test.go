package problem942

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type testUnit struct {
	input  string
	output []int
}

func TestDiStringMatch(t *testing.T) {
	table := []testUnit{
		{
			"IDID",
			[]int{0, 4, 1, 3, 2},
		},
		{
			"III",
			[]int{0, 1, 2, 3},
		},
		{
			"DDI",
			[]int{3, 2, 0, 1},
		},
	}

	for _, tt := range table {
		if actual := diStringMatch(tt.input); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("diStringMatch(%s) expects %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
