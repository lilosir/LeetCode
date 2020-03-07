package problem905

// func sortArrayByParity(A []int) []int {
// 	result := make([]int, len(A))
// 	n := 0
// 	m := 0
// 	for i := 0; i < len(A); i++ {
// 		if A[i]%2 == 0 {
// 			result[n] = A[i]
// 			n++
// 		} else {
// 			result[len(A)-1-m] = A[i]
// 			m++
// 		}
// 	}

// 	return result
// }

func sortArrayByParity(A []int) []int {
	for i, j := 0, 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[i], A[j] = A[j], A[i]
			j++
		}
	}

	return A
}
