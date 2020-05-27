package problem1046

import "testing"

type testUnit struct {
	input  []int
	output int
}

func TestLastStoneWeight(t *testing.T) {
	table := []testUnit{
		{
			[]int{2, 7, 4, 1, 8, 1},
			1,
		},
		{
			[]int{2, 7},
			5,
		},
		{
			[]int{2, 2},
			0,
		},
	}

	for _, tt := range table {
		if actual := lastStoneWeight(tt.input); actual != tt.output {
			t.Errorf("lastStoneWeight(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
