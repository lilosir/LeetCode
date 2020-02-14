package problem1342

func numberOfSteps(num int) int {
	result := 0
	for num > 0 {
		if num%2 == 1 {
			num--
		} else {
			num = num >> 1
		}
		result++
	}
	return result
}
