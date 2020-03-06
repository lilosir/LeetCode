package problem832

import (
	"testing"
)

type testUnit struct {
	input  [][]int
	output [][]int
}

func TestFlipAndInvertImage(t *testing.T) {
	table := []testUnit{
		{
			[][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			[][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		{
			[][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
			[][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
				{1, 0, 1, 0},
			},
		},
	}

	for _, tt := range table {
		if actual := flipAndInvertImage(tt.input); !compareTwoDimensionSlices(actual, tt.output) {
			t.Errorf("flipAndInvertImage(%v) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}

func compareTwoDimensionSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if len(v) != len(b[i]) {
			return false
		}
		for j, val := range v {
			if val != b[i][j] {
				return false
			}
		}
	}
	return true
}
