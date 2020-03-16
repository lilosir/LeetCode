package problem1380

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type testUnit struct {
	input  [][]int
	output []int
}

func TestLuckyNumbers(t *testing.T) {
	table := []testUnit{
		// {
		// 	[][]int{
		// 		{
		// 			3, 7, 8,
		// 		},
		// 		{
		// 			9, 11, 13,
		// 		},
		// 		{
		// 			15, 16, 17,
		// 		},
		// 	},
		// 	[]int{15},
		// },
		// {
		// 	[][]int{
		// 		{
		// 			1, 10, 4, 2,
		// 		},
		// 		{
		// 			9, 3, 8, 7,
		// 		},
		// 		{
		// 			15, 16, 17, 12,
		// 		},
		// 	},
		// 	[]int{12},
		// },
		// {
		// 	[][]int{
		// 		{
		// 			7, 8,
		// 		},
		// 		{
		// 			1, 2,
		// 		},
		// 	},
		// 	[]int{7},
		// },
		{
			[][]int{
				{
					41205, 87385,
				},
				{
					71957, 59065,
				},
				{
					57152, 61601,
				},
			},
			[]int{},
		},
	}

	for _, tt := range table {
		if actual := luckyNumbers(tt.input); !kit.CompareTwoSlice(actual, tt.output) {
			t.Errorf("luckyNumbers(%v) expects %v, but got %v", tt.input, tt.output, actual)
		}
	}
}
