package problem999

func numRookCaptures(board [][]byte) int {
	res := 0
	x, y := 0, 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				x, y = i, j
				break
			}
		}
	}

	for i := x; i > 0; i-- {
		if board[i][y] == 'p' {
			res++
			break
		} else if board[i][y] == 'B' {
			break
		}
	}

	for i := x; i < 8; i++ {
		if board[i][y] == 'p' {
			res++
			break
		} else if board[i][y] == 'B' {
			break
		}
	}

	for i := y; i > 0; i-- {
		if board[x][i] == 'p' {
			res++
			break
		} else if board[x][i] == 'B' {
			break
		}
	}

	for i := y; i < 8; i++ {
		if board[x][i] == 'p' {
			res++
			break
		} else if board[x][i] == 'B' {
			break
		}
	}

	return res
}
