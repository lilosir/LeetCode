package problem832

// func flipAndInvertImage(A [][]int) [][]int {
// 	result := make([][]int, len(A))
// 	insideLen := len(A[0])
// 	for i := range result {
// 		result[i] = make([]int, insideLen)
// 	}

// 	for j := range A {
// 		for k := range result[j] {
// 			if A[j][k] > 0 {
// 				result[j][insideLen-1-k] = 0
// 			} else {
// 				result[j][insideLen-1-k] = 1
// 			}
// 		}
// 	}

// 	return result
// }

// Bitwise XOR !!!
// func flipAndInvertImage(A [][]int) [][]int {
// 	result := make([][]int, len(A))
// 	insideLen := len(A[0])
// 	for i := range result {
// 		result[i] = make([]int, insideLen)
// 	}

// 	for j := range A {
// 		for k := range result[j] {
// 			result[j][insideLen-1-k] = A[j][k] ^ 1
// 		}
// 	}

// 	return result
// }

func flipAndInvertImage(A [][]int) [][]int {

	for _, a := range A {
		for j := 0; j <= (len(a)-1)/2; j++ {
			a[j], a[len(a)-1-j] = a[len(a)-1-j]^1, a[j]^1
		}
	}

	return A
}
