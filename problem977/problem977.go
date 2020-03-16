package problem977

func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	left, right := 0, len(A)-1
	for i := len(A) - 1; i >= 0; i-- {
		pow2Left, pow2Right := A[left]*A[left], A[right]*A[right]
		if pow2Left > pow2Right {
			result[i] = pow2Left
			left++
		} else {
			result[i] = pow2Right
			right--
		}
	}
	return result
}
