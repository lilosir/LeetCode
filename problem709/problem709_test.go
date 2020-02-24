package problem709

import "testing"

type testUnit struct {
	input  string
	output string
}

func TestToLowerCase(t *testing.T) {
	table := []testUnit{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
	}

	for _, tt := range table {
		if actual := toLowerCase(tt.input); actual != tt.output {
			t.Errorf("toLowerCase(%s) expect %s, but got %s", tt.input, tt.output, actual)
		}
	}
}
