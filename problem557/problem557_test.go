package problem557

import (
	"testing"
)

type testUnit struct {
	input  string
	output string
}

func TestReverseWords(t *testing.T) {
	table := []testUnit{
		{
			"Let's take LeetCode contest",
			"s'teL ekat edoCteeL tsetnoc",
		},
	}

	for _, tt := range table {
		if actual := reverseWords(tt.input); actual != tt.output {
			t.Errorf("reverseWords(%s) expect \"%s\", but got \"%s\"", tt.input, tt.output, actual)
		}
	}
}
