package problem922

// func sortArrayByParityII(A []int) []int {
// 	for i := 0; i < len(A)-1; i++ {
// 		if i&1 == 0 && A[i]&1 == 1 {
// 			for j := i + 1; j < len(A); j += 2 {
// 				if A[j]&1 == 0 {
// 					A[i], A[j] = A[j], A[i]
// 				}
// 			}
// 		}
// 		if i&1 == 1 && A[i]&1 == 0 {
// 			for j := i + 1; j < len(A); j += 2 {
// 				if A[j]&1 == 1 {
// 					A[i], A[j] = A[j], A[i]
// 				}
// 			}
// 		}
// 	}
// 	return A
// }

func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A)-1; i += 2 {
		if A[i]&1 == 1 {
			for A[j]&1 == 1 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
