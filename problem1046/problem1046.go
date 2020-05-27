package problem1046

import (
	"sort"
)

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})

	for len(stones) > 1 {
		if stones[0] > stones[1] {
			stones = append(stones[2:], stones[0]-stones[1])
		} else {
			stones = stones[2:]
		}
		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
