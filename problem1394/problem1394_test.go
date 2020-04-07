package problem1394

import "testing"

type testUnit struct {
	input  []int
	output int
}

func TestFindLucky(t *testing.T) {
	table := []testUnit{
		{
			[]int{2, 2, 3, 4},
			2,
		},
		{
			[]int{1, 2, 2, 3, 3, 3},
			3,
		},
		{
			[]int{2, 2, 2, 3, 3},
			-1,
		},
		{
			[]int{7, 7, 7, 7, 7, 7, 7},
			7,
		},
	}

	for _, tt := range table {
		if actual := findLucky(tt.input); actual != tt.output {
			t.Errorf("findLucky(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
