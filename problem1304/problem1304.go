package problem1304

func sumZero(n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		// mod := n % 2
		// if mod == 0 {
		// 	result[i] = i-n/2
		// } else {
		// 	result[i] = i-n/2
		// }
		result[i] = i - n/2
	}
	return result
}
