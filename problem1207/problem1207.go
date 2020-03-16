package problem1207

func uniqueOccurrences(arr []int) bool {
	numMap := make(map[int]int)

	for _, n := range arr {
		numMap[n]++
	}

	lenMap := make(map[int]int)

	for _, v := range numMap {
		if _, ok := lenMap[v]; ok {
			return false
		}
		lenMap[v] = 0
	}

	return true
}
