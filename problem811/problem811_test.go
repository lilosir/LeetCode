package problem811

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type testUnit struct {
	input  []string
	output []string
}

func TestSubdomainVisits(t *testing.T) {
	table := []testUnit{
		{
			[]string{"9001 discuss.leetcode.com"},
			[]string{
				"9001 discuss.leetcode.com",
				"9001 leetcode.com",
				"9001 com",
			},
		},
		{
			[]string{
				"900 google.mail.com",
				"50 yahoo.com",
				"1 intel.mail.com",
				"5 wiki.org",
			},
			[]string{
				"901 mail.com",
				"50 yahoo.com",
				"900 google.mail.com",
				"5 wiki.org",
				"5 org",
				"1 intel.mail.com",
				"951 com",
			},
		},
	}

	for _, tt := range table {
		actual := subdomainVisits(tt.input)
		actualInterface := make([]interface{}, len(actual))
		for _, v := range actual {
			actualInterface = append(actualInterface, v)
		}
		outputInterface := make([]interface{}, len(tt.output))
		for _, v := range tt.output {
			outputInterface = append(outputInterface, v)
		}
		if !kit.CheckTwoSliceHaveSameElements(actualInterface, outputInterface) {
			t.Errorf("subdomainVisits(%v) expects %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
