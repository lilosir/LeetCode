package problem1299

func replaceElements(arr []int) []int {
	if len(arr) == 1 {
		return []int{-1}
	}

	result := make([]int, len(arr))
	result[len(arr)-1] = -1
	result[len(arr)-2] = arr[len(arr)-1]

	for i := len(arr) - 2; i > 0; i-- {
		if arr[i] > result[i] {
			result[i-1] = arr[i]
		} else {
			result[i-1] = result[i]
		}
	}
	return result
}

// good solution

// func replaceElements(arr []int) []int {
//     mx := -1
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		arr[i], mx = mx, int(math.Max(float64(mx), float64(arr[i])))
// 	}
// 	return arr
// }
