package problem883

import "testing"

type testUnit struct {
	input  [][]int
	output int
}

func TestProjectionArea(t *testing.T) {
	table := []testUnit{
		{
			[][]int{
				{2},
			},
			5,
		},
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			17,
		},
		{
			[][]int{
				{1, 0},
				{0, 2},
			},
			8,
		},
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			14,
		},
		{
			[][]int{
				{2, 2, 2},
				{2, 1, 2},
				{2, 2, 2},
			},
			21,
		},
	}

	for _, tt := range table {
		if actual := projectionArea(tt.input); actual != tt.output {
			t.Errorf("projectionArea(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
