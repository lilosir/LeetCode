package problem961

import "testing"

type testUnit struct {
	input  []int
	output int
}

func TestRepeatedNTimes(t *testing.T) {
	table := []testUnit{
		{
			[]int{1, 2, 3, 3},
			3,
		},
		{
			[]int{2, 1, 2, 5, 3, 2},
			2,
		},
		{
			[]int{5, 1, 5, 2, 5, 3, 5, 4},
			5,
		},
		{
			[]int{9, 5, 6, 9},
			9,
		},
	}

	for _, tt := range table {
		if actual := repeatedNTimes(tt.input); actual != tt.output {
			t.Errorf("repeatedNTimes(%v), expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
