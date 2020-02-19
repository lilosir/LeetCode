package problem1221

import (
	"testing"
)

type testUnit struct {
	input  string
	output int
}

func TestBalancedStringSplit(t *testing.T) {
	table := []testUnit{
		{"RLRRLLRLRL", 4},
		{"RLLLLRRRLR", 3},
		{"LLLLRRRR", 1},
		{"RLRRRLLRLL", 2},
	}

	for _, tt := range table {
		if actual := balancedStringSplit(tt.input); actual != tt.output {
			t.Errorf("balancedStringSplit(%s) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
