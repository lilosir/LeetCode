package problem1295

// func findNumbers(nums []int) int {
// 	total := 0
// 	for _, v := range nums {
// 		len := 0
// 		for v > 0 {
// 			len++
// 			v = v / 10
// 		}
// 		if len%2 == 0 {
// 			total++
// 		}
// 	}
// 	return total
// }

func findNumbers(nums []int) int {
	total := 0
	for _, v := range nums {
		if (v > 9 && v < 100) || (v > 999 && v < 10000) {
			total++
		}
	}
	return total
}
