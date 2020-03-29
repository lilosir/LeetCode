package problem1122

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	countMap := make(map[int]int)
	ascendTail := []int{}
	for _, v := range arr2 {
		countMap[v] = 0
	}

	for _, v1 := range arr1 {
		if _, ok := countMap[v1]; !ok {
			ascendTail = append(ascendTail, v1)
		} else {
			countMap[v1]++
		}
	}

	res := make([]int, 0, len(arr1))
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < countMap[arr2[i]]; j++ {
			res = append(res, arr2[i])
		}
	}

	sort.Ints(ascendTail)
	res = append(res, ascendTail...)

	return res
}
