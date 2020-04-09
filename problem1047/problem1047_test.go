package problem1047

import (
	"testing"
)

type testUnit struct {
	input  string
	output string
}

func TestRemoveDuplicates(t *testing.T) {
	table := []testUnit{
		{
			"abbaca",
			"ca",
		},
	}

	for _, tt := range table {
		if actual := removeDuplicates(tt.input); actual != tt.output {
			t.Errorf("removeDuplicates(%s) expect %s, but got %s", tt.input, tt.output, actual)
		}
	}
}
