package problem1252

import "testing"

type testUnit struct {
	input  inputStruct
	output int
}

type inputStruct struct {
	n       int
	m       int
	indices [][]int
}

func TestOddCells(t *testing.T) {
	table := []testUnit{
		{
			inputStruct{
				2,
				3,
				[][]int{
					{0, 1},
					{1, 1},
				},
			},
			6,
		},
		{
			inputStruct{
				2,
				2,
				[][]int{
					{1, 1},
					{0, 0},
				},
			},
			0,
		},
	}

	for _, tt := range table {
		if actual := oddCells(tt.input.n, tt.input.m, tt.input.indices); actual != tt.output {
			t.Errorf("toLowerCase(%d, %d, %v) expect %d, but got %d", tt.input.n, tt.input.m, tt.input.indices, tt.output, actual)
		}
	}
}
