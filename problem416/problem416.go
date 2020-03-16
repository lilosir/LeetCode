package problem416

func hammingDistance(x int, y int) int {
	xorResult := x ^ y
	result := 0
	for xorResult > 0 {
		if xorResult&1 == 1 {
			result++
		}
		xorResult >>= 1
	}
	return result
}
