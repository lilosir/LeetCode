package problem1313

import "testing"

type testUnit struct {
	input  []int
	output []int
}

func TestDecompressRLElist(t *testing.T) {
	table := []testUnit{
		{[]int{1, 2, 3, 4}, []int{2, 4, 4, 4}},
	}

	for _, tt := range table {
		if actual := decompressRLElist(tt.input); !compareTwoArrays(actual, tt.output) {
			t.Errorf("decompressRLElist(%v) expect %v, but got %v", tt.input, tt.output, actual)
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
