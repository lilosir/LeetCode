package problem1108

import (
	"testing"
)

type testTable struct {
	input string
	output string
}

func TestDefangIPaddr(t *testing.T) {
	table := []testTable{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
		{"192.168.2.23", "192[.]168[.]2[.]23"},
	}

	for _, tt := range table {
		if actual := defangIPaddr(tt.input); actual != tt.output {
			t.Errorf("defangIPaddr(%s) expect %s, but got %s", tt.input, tt.output, actual)
		}
	}

}