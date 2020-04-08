package problem1002

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type testUnit struct {
	input  []string
	output []string
}

func TestCommonChars(t *testing.T) {
	table := []testUnit{
		{
			[]string{"bella", "label", "roller"},
			[]string{"e", "l", "l"},
		},
		{
			[]string{"cool", "lock", "cook"},
			[]string{"c", "o"},
		},
	}

	for _, tt := range table {
		if actual := commonChars(tt.input); !kit.CompareTwoStringSlice(actual, tt.output) {
			t.Errorf("commonChars(%v) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
