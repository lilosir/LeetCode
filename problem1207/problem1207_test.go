package problem1207

import "testing"

type testUnit struct {
	input  []int
	output bool
}

func TestUniqueOccurrences(t *testing.T) {
	table := []testUnit{
		{
			[]int{1, 2, 2, 1, 1, 3},
			true,
		},
		{
			[]int{1, 2},
			false,
		},
		{
			[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			true,
		},
	}

	for _, tt := range table {
		if actual := uniqueOccurrences(tt.input); actual != tt.output {
			t.Errorf("uniqueOccurrences(%v) expect %t, but got %t", tt.input, tt.output, actual)
		}
	}
}
