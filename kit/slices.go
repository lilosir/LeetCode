package kit

// CompareTwoSlice return two int sclices are same
func CompareTwoSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// CheckTwoSliceHaveSameStringElements return true if they have same elements--- type string
func CheckTwoSliceHaveSameStringElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	valMap := make(map[string]int)

	for _, v := range a {
		valMap[v]++
	}

	for _, v := range b {
		if _, ok := valMap[v]; !ok {
			return false
		}
		valMap[v]--
		if valMap[v] == 0 {
			delete(valMap, v)
		}
	}

	if len(valMap) == 0 {
		return true
	}

	return false
}
