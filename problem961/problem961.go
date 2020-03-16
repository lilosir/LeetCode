package problem961

// func repeatedNTimes(A []int) int {
// 	count := make(map[int]int)

// 	for _, v := range A {
// 		if val, ok := count[v]; !ok {
// 			count[v] = 1
// 		} else {
// 			count[v] = val + 1
// 			if count[v] == len(A)/2 {
// 				return v
// 			}
// 		}
// 	}
// 	return 0
// }

func repeatedNTimes(A []int) int {
	for i := 2; i < len(A); i++ {
		if A[i-1] == A[i] || A[i-2] == A[i] {
			return A[i]
		}
	}
	return A[0]
}
