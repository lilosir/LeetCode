package problem771

import (
	"testing"
)

type testUnit struct {
	input  []string
	output int
}

func TestNumJewelsInStones(t *testing.T) {
	table := []testUnit{
		{
			[]string{"aA", "aAAbbbb"},
			3,
		},
		{
			[]string{"z", "ZZ"},
			0,
		},
	}

	for _, tt := range table {
		if actual := numJewelsInStones(tt.input[0], tt.input[1]); actual != tt.output {
			t.Errorf("numJewelsInStones(%s, %s), expect %d, but got %d", tt.input[0], tt.input[1], tt.output, actual)
		}
	}
}
