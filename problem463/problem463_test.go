package problem463

import (
	"testing"
)

type testUnit struct {
	input  [][]int
	output int
}

func TestIslandPerimeter(t *testing.T) {
	table := []testUnit{
		{
			[][]int{
				{
					0, 1, 0, 0,
				},
				{
					1, 1, 1, 0,
				},
				{
					0, 1, 0, 0,
				},
				{
					1, 1, 0, 0,
				},
			},
			16,
		},
	}

	for _, tt := range table {
		if actual := islandPerimeter(tt.input); actual != tt.output {
			t.Errorf("islandPerimeter(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
