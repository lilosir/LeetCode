package problem883

func projectionArea(grid [][]int) int {
	x, y, z := 0, 0, 0
	for i := range grid[0] {
		xtemp, ytemp := 0, 0
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > xtemp {
				xtemp = grid[i][j]
			}
			if grid[j][i] > ytemp {
				ytemp = grid[j][i]
			}
			if grid[i][j] > 0 {
				z++
			}
		}
		x += xtemp
		y += ytemp
	}
	return x + y + z
}
