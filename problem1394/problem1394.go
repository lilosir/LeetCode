package problem1394

func findLucky(arr []int) int {
	res := make([]int, 501)

	for _, v := range arr {
		res[v]++
	}

	for i := len(res) - 1; i >= 0; i-- {
		if i == res[i] {
			return i
		}
	}
	return -1
}
