package problem1365

import "testing"

type testUnit struct {
	input  []int
	output []int
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	table := []testUnit{
		{
			[]int{8, 1, 2, 2, 3},
			[]int{4, 0, 1, 1, 3},
		},
		{
			[]int{6, 5, 4, 8},
			[]int{2, 1, 0, 3},
		},
		{
			[]int{7, 7, 7, 7},
			[]int{0, 0, 0, 0},
		},
	}

	for _, tt := range table {
		if actual := smallerNumbersThanCurrent(tt.input); !compareTwoArrays(actual, tt.output) {
			t.Errorf("smallerNumbersThanCurrent(%v) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}

func compareTwoArrays(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, key := range a {
		if key != b[i] {
			return false
		}
	}
	return true
}
