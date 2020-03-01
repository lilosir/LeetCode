package problem1299

import "testing"

type testUnit struct {
	input  []int
	output []int
}

func TestReplaceElements(t *testing.T) {
	table := []testUnit{
		{
			[]int{17, 18, 5, 4, 6, 1},
			[]int{18, 6, 6, 6, 1, -1},
		},
		{
			[]int{22, 8, 13, 55, 3, 11},
			[]int{55, 55, 55, 11, 11, -1},
		},
		{
			[]int{400},
			[]int{-1},
		},
	}

	for _, tt := range table {
		if actual := replaceElements(tt.input); !compareTwoSlice(actual, tt.output) {
			t.Errorf("replaceElements(%v) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}

func compareTwoSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
