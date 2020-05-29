package problem463

func islandPerimeter(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				result += checkAround(i, j, grid)
			}
		}
	}
	return result
}

func checkAround(x, y int, grid [][]int) int {
	leng := 0
	//up
	if x > 0 && grid[x-1][y] == 0 || x == 0 {
		leng++
	}
	//down
	if x < len(grid)-1 && grid[x+1][y] == 0 || x == len(grid)-1 {
		leng++
	}
	//left
	if y > 0 && grid[x][y-1] == 0 || y == 0 {
		leng++
	}
	//right
	if y < len(grid[x])-1 && grid[x][y+1] == 0 || y == len(grid[x])-1 {
		leng++
	}
	return leng
}
