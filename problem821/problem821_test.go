package problem821

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type input struct {
	a string
	b byte
}

type testUnit struct {
	input  input
	output []int
}

func TestSumZero(t *testing.T) {
	table := []testUnit{
		{
			input{
				"loveleetcode",
				'e',
			},
			[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			input{
				"aaba",
				'b',
			},
			[]int{2, 1, 0, 1},
		},
	}

	for _, tt := range table {
		if actual := shortestToChar(tt.input.a, tt.input.b); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("shortestToChar(%s %x) expect %v, but got %v", tt.input.a, tt.input.b, tt.output, actual)
		}
	}
}
