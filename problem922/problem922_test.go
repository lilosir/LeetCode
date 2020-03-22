package problem922

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type testUnit struct {
	input  []int
	output []int
}

func TestSortArrayByParityII(t *testing.T) {
	table := []testUnit{
		{
			[]int{4, 2, 5, 7},
			[]int{4, 5, 2, 7},
		},
		{
			[]int{1, 10, 5, 6, 7, 8},
			[]int{10, 1, 6, 5, 8, 7},
		},
	}

	for _, tt := range table {
		actual := sortArrayByParityII(tt.input)
		actualInterface := make([]interface{}, len(actual))
		for _, v := range actual {
			actualInterface = append(actualInterface, v)
		}
		outputInterface := make([]interface{}, len(tt.output))
		for _, v := range tt.output {
			outputInterface = append(outputInterface, v)
		}
		if !kit.CheckTwoSliceHaveSameElements(actualInterface, outputInterface) || !checkPositionValidation(actual) {
			t.Errorf("sortArrayByParityII(%v) expect %v, but got %v", tt.input, tt.output, actual)
		}
	}
}

func checkPositionValidation(a []int) bool {
	for i, v := range a {
		if (i%2 == 0 && v%2 == 1) || (i%2 == 1 && v%2 == 0) {
			return false
		}
	}
	return true
}
