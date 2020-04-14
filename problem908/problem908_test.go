package problem908

import "testing"

type input struct {
	a []int
	b int
}

type testUnit struct {
	input  input
	output int
}

func TestSmallestRangeI(t *testing.T) {
	table := []testUnit{
		{
			input{
				a: []int{1},
				b: 0,
			},
			0,
		},
		{
			input{
				a: []int{0, 10},
				b: 2,
			},
			6,
		},
		{
			input{
				a: []int{1, 3, 6},
				b: 3,
			},
			0,
		},
	}

	for _, tt := range table {
		if actual := smallestRangeI(tt.input.a, tt.input.b); actual != tt.output {
			t.Errorf("smallestRangeI(%v, %d) expect %v, but got %v", tt.input.a, tt.input.b, tt.output, actual)
		}
	}
}
