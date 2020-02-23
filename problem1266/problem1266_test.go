package problem1266

import "testing"

type testUnit struct {
	input  [][]int
	output int
}

func TestMinTimeToVisitAllPoints(t *testing.T) {
	table := []testUnit{
		{
			[][]int{
				{1, 1},
				{3, 4},
				{-1, 0},
			},
			7,
		},
		{
			[][]int{
				{3, 2},
				{-2, 2},
			},
			5,
		},
	}

	for _, tt := range table {
		if actual := minTimeToVisitAllPoints(tt.input); actual != tt.output {
			t.Errorf("subtractProductAndSum(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
