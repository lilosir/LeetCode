package problem852

import "testing"

type testUnit struct {
	input  []int
	output int
}

func TestPeakIndexInMountainArray(t *testing.T) {
	table := []testUnit{
		{
			[]int{0, 1, 0},
			1,
		},
		{
			[]int{0, 2, 1, 0},
			1,
		},
	}
	for _, tt := range table {
		if actual := peakIndexInMountainArray(tt.input); actual != tt.output {
			t.Errorf("peakIndexInMountainArray(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
