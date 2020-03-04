package problem1304

func sumZero(n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		if n%2 == 0 {
			result[i] = i - n/2
			if i >= n/2 {
				result[i] = i - n/2 + 1
			}
		} else {
			result[i] = i - n/2
		}
	}
	return result
}
