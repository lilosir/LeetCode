package problem1051

import (
	"sort"
)

func heightChecker(heights []int) int {
	result := 0
	oriHeights := make([]int, len(heights))
	copy(oriHeights, heights)

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})

	for i := range heights {
		if heights[i] != oriHeights[i] {
			result++
		}
	}
	return result
}
