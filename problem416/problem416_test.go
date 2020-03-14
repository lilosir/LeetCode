package problem416

import (
	"testing"
)

type testUnit struct {
	input  []int
	output int
}

func TestHammingDistance(t *testing.T) {
	table := []testUnit{
		{
			[]int{1, 4},
			2,
		},
	}

	for _, tt := range table {
		if actual := hammingDistance(tt.input[0], tt.input[1]); actual != tt.output {
			t.Errorf("hammingDistance(%d, %d), expect %d, but got %d", tt.input[0], tt.input[1], tt.output, actual)
		}
	}
}
