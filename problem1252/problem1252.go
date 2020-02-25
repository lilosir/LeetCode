package problem1252

func oddCells(n int, m int, indices [][]int) int {
	result := 0
	initialMatrix := make([][]int, n)
	for i := range initialMatrix {
		initialMatrix[i] = make([]int, m)
	}

	for i := 0; i < len(indices); i++ {
		r := indices[i][0]
		c := indices[i][1]
		for j := 0; j < m; j++ {
			initialMatrix[r][j]++
			if initialMatrix[r][j]%2 == 1 {
				result++
			} else {
				result--
			}
		}
		for k := 0; k < n; k++ {
			initialMatrix[k][c]++
			if initialMatrix[k][c]%2 == 1 {
				result++
			} else {
				result--
			}
		}
	}

	return result
}
