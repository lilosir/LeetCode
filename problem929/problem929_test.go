package problem929

import (
	"testing"
)

type testUnit struct {
	input  []string
	output int
}

func TestNumUniqueEmails(t *testing.T) {
	table := []testUnit{
		{
			[]string{
				"test.email+alex@leetcode.com",
				"test.e.mail+bob.cathy@leetcode.com",
				"testemail+david@lee.tcode.com",
			},
			2,
		},
		{
			[]string{
				"test.email+alex@leetcode.com",
				"test.email.leet+alex@code.com",
			},
			2,
		},
	}

	for _, tt := range table {
		if actual := numUniqueEmails(tt.input); actual != tt.output {
			t.Errorf("numUniqueEmails(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
