package problem1365

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	copyNum := make([]int, len(nums))
	copy(copyNum, nums)
	sort.Ints(nums)
	mapNum := make(map[int]int)
	for i, v := range nums {
		_, ok := mapNum[v]
		if !ok {
			mapNum[v] = i
		}
	}

	for i := 0; i < len(copyNum); i++ {
		result[i] = mapNum[copyNum[i]]
	}
	return result
}
