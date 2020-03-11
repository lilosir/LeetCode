package problem728

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ {
		if getSelfDivisible(i) {
			result = append(result, i)
		}
	}
	return result
}

func getSelfDivisible(num int) bool {
	if num > 0 && num < 10 {
		return true
	}
	og := num
	for num > 0 {
		if num%10 == 0 {
			return false
		}
		temp := num
		num /= 10

		if og%(temp-num*10) != 0 {
			return false
		}
	}
	return true
}
