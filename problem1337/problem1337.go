package problem1337

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	result := make([]int, k)
	newRow := make([][2]int, len(mat))

	for i := 0; i < len(mat); i++ {
		power := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				power++
			} else {
				break
			}
		}
		newRow[i] = [2]int{i, power}
	}

	sort.Slice(newRow, func(a, b int) bool {
		if newRow[a][1] == newRow[b][1] {
			return newRow[a][0] < newRow[b][0]
		}
		return newRow[a][1] < newRow[b][1]
	})

	for i := 0; i < k; i++ {
		result[i] = newRow[i][0]
	}
	return result
}
