package problem665

import "testing"

type testUnit struct {
	input  []int
	output bool
}

func TestCheckPossibility(t *testing.T) {
	table := []testUnit{
		{
			[]int{4, 2, 3},
			true,
		},
		{
			[]int{4, 2, 1},
			false,
		},
		{
			[]int{1, 4, 2, 3},
			true,
		},
		{
			[]int{1, 4, 2, 3, 2},
			false,
		},
		// {
		// 	[]int{2, 3, 3, 2, 4},
		// 	true,
		// },
		// {
		// 	[]int{3, 4, 2, 3},
		// 	false,
		// },
		// {
		// 	[]int{1, 3, 2},
		// 	true,
		// },
	}

	for _, tt := range table {
		if actual := checkPossibility(tt.input); actual != tt.output {
			t.Errorf("checkPossibility(%v) expect %t, but got %t", tt.input, tt.output, actual)
		}
	}
}
