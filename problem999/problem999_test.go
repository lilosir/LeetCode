package problem999

import (
	"testing"
)

type testUnit struct {
	input  [][]byte
	output int
}

func TestNumRookCaptures(t *testing.T) {
	table := []testUnit{
		{
			[][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', 'R', '.', '.', '.', 'p'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			},
			3,
		},
		{
			[][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
				{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
				{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
				{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
				{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			},
			0,
		},
	}

	for _, tt := range table {
		if actual := numRookCaptures(tt.input); actual != tt.output {
			t.Errorf("numRookCaptures(%v) expects %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
