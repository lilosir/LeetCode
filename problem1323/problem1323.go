package problem1323

import (
	"math"
)

// func maximum69Number(num int) int {
// 	numStr := strconv.Itoa(num)
// 	max := num
// 	for k, v := range numStr {
// 		current := num
// 		if v == '6' {
// 			current += int(math.Pow10(len(numStr)-k-1)) * 3
// 		}
// 		if current > max {
// 			max = current
// 		}
// 	}
// 	return max
// }

func maximum69Number(num int) int {
	len := 0
	newNum := num
	max := num
	for newNum > 0 {
		current := num
		if newNum%10 == 6 {
			current += int(math.Pow10(len) * 3)
		}
		if max < current {
			max = current
		}
		newNum /= 10
		len++
	}

	return max
}
