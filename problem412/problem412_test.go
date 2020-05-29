package problem412

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type testUnit struct {
	input  int
	output []string
}

func TestFizzBuzz(t *testing.T) {
	table := []testUnit{
		{
			15,
			[]string{"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}

	for _, tt := range table {
		if actual := fizzBuzz(tt.input); !kit.CompareTwoStringSlice(actual, tt.output) {
			t.Errorf("fizzBuzz(%d) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
