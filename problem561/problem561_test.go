package problem561

import "testing"

type testUnit struct {
	input  []int
	output int
}

func TestArrayPairSum(t *testing.T) {
	table := []testUnit{
		{
			[]int{1,4,3,2}, 
			4,
		},
	}
	for _, tt := range table {
		if actual := arrayPairSum(tt.input); actual != tt.output {
			t.Errorf("arrayPairSum(%v) expect %v, but got %d", tt.input, tt.output, actual)
		}
	}
}
