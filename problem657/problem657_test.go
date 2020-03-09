package problem657

import "testing"

type testUnit struct {
	input  string
	output bool
}

func TestJudgeCircle(t *testing.T) {
	table := []testUnit{
		{
			"UD",
			true,
		},
		{
			"LL",
			false,
		},
		{
			"LUUDD",
			false,
		},
		{
			"LUUDDR",
			true,
		},
	}

	for _, tt := range table {
		if actual := judgeCircle(tt.input); actual != tt.output {
			t.Errorf("judgeCircle(%s) expect %t, but got %t", tt.input, tt.output, actual)
		}
	}
}
