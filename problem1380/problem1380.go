package problem1380

func luckyNumbers(matrix [][]int) []int {
	result := []int{}
	minMap := make(map[int]int)

	for i := 0; i < len(matrix); i++ {
		min := 100000
		j := 0
		for k := 0; k < len(matrix[i]); k++ {
			if matrix[i][k] < min {
				min = matrix[i][k]
				j = k
			}
		}
		minMap[i] = j
	}

	for i := 0; i < len(matrix[0]); i++ {
		max := 0
		j := 0
		for k := 0; k < len(matrix); k++ {
			if matrix[k][i] > max {
				max = matrix[k][i]
				j = k
			}
		}
		maxMap[i] = j

		if v, ok := minMap[j]; ok && v == i {
			result = append(result, matrix[j][i])
		}
	}
	return result
}
