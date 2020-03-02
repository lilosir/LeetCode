package problem1021

import "testing"

type testUnit struct {
	input  string
	output string
}

func TestRemoveOuterParentheses(t *testing.T) {
	table := []testUnit{
		{
			"(()())(())",
			"()()()",
		},
		{
			"(()())(())(()(()))",
			"()()()()(())",
		},
		{
			"()()",
			"",
		},
	}

	for _, tt := range table {
		if actual := removeOuterParentheses(tt.input); actual != tt.output {
			t.Errorf("removeOuterParentheses(%s) expect %s, but got %s", tt.input, tt.output, actual)
		}
	}
}
