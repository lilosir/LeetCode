package problem1295

import "testing"

type testUnit struct {
	input  []int
	output int
}

func TestFindNumbers(t *testing.T) {
	table := []testUnit{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}

	for _, tt := range table {
		if actual := findNumbers(tt.input); actual != tt.output {
			t.Errorf("subtractProductAndSum(%d) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
