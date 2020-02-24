package problem1323

import "testing"

type testUnit struct {
	input  int
	output int
}

func TestMaximum69Number(t *testing.T) {

	table := []testUnit{
		{
			9669,
			9969,
		},
		{
			9996,
			9999,
		},
		{
			9999,
			9999,
		},
	}
	for _, tt := range table {
		if actual := maximum69Number(tt.input); actual != tt.output {
			t.Errorf("miaximum69Number(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
