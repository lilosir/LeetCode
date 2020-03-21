package problem1337

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type inputType struct {
	a [][]int
	b int
}

type testUnit struct {
	input  inputType
	output []int
}

func TestKWeakestRows(t *testing.T) {
	table := []testUnit{
		{
			inputType{
				[][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 0},
					{1, 0, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 1},
				},
				3,
			},
			[]int{2, 0, 3},
		},
		{
			inputType{
				[][]int{
					{1, 0, 0, 0},
					{1, 1, 1, 1},
					{1, 0, 0, 0},
					{1, 0, 0, 0},
				},
				2,
			},
			[]int{0, 2},
		},
		{
			inputType{
				[][]int{
					{1, 1, 0},
					{1, 1, 0},
					{1, 1, 1},
					{1, 1, 1},
					{0, 0, 0},
					{1, 1, 1},
					{1, 0, 0},
				},
				6,
			},
			[]int{4, 6, 0, 1, 2, 3},
		},
	}
	for _, tt := range table {
		if actual := kWeakestRows(tt.input.a, tt.input.b); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("kWeakestRows(%v) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
