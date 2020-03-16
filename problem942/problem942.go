package problem942

func diStringMatch(S string) []int {
	result := make([]int, len(S)+1)

	left, right := 0, len(S)
	for i, s := range S {
		if s == 'I' {
			result[i] = left
			left++
		} else {
			result[i] = right
			right--
		}
	}
	result[len(S)] = left
	return result
}
