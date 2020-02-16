package problem1281

import (
	"testing"
)

type testUnit struct {
	input  int
	output int
}

func TestSubtractProductAndSum(t *testing.T) {
	table := []testUnit{
		{234, 15},
		{4421, 21},
	}

	for _, tt := range table {
		if actual := subtractProductAndSum(tt.input); actual != tt.output {
			t.Errorf("subtractProductAndSum(%d) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
