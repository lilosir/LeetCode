package problem1351

import "testing"

type testUnit struct {
	input  [][]int
	output int
}

func TestCountNegatives(t *testing.T) {
	table := []testUnit{
		{
			[][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			8,
		},
		{
			[][]int{
				{3, 2},
				{1, 0},
			},
			0,
		},
		{
			[][]int{
				{1, -1},
				{-1, -1},
			},
			3,
		},
		{
			[][]int{
				{-1},
			},
			1,
		},
	}

	for _, tt := range table {
		if actual := countNegatives(tt.input); actual != tt.output {
			t.Errorf("countNegatives(%v) expect %d, but got %d\n", tt.input, tt.output, actual)
		}
	}
}
