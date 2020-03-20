package problem1356

import (
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(a, b int) bool {
		c := helper(arr[a])
		d := helper(arr[b])
		if c == d {
			return arr[a] < arr[b]
		}
		return c < d
	})
	return arr
}

func helper(n int) int {
	c := 0
	for n != 0 {
		c += n & 1
		n >>= 1
	}
	return c
}
