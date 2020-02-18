package problem1351

// 16 ms, 6MB
// func countNegatives(grid [][]int) int {
//     result := 0
// 	n := len(grid[0])
// 	for _, row := range grid {
// 		for j, num := range row {
// 			if num < 0 {
// 				totalInRow := n - j
// 				result += totalInRow
// 				break
// 			}
// 		}
// 	}

// 	return result
// }

// 12 ms, 6MB
func countNegatives(grid [][]int) int {
	result := 0
	m := len(grid)
	n := len(grid[0])
	current := 0
	for m > 0 && current < n {
		if grid[m-1][n-1] >= 0 {
			current = 0
			m--
		} else if grid[m-1][current] < 0 {
			result += n - current
			m--
		} else {
			current++
		}
	}

	return result
}
