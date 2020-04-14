package problem908

func smallestRangeI(A []int, K int) int {
	min := A[0]
	max := A[0]
	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
		if A[i] < min {
			min = A[i]
		}
	}
	if max-min <= K*2 {
		return 0
	}

	return max - min - 2*K
}
