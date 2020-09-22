func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			processCell(board, i, j)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] > 2 {
				board[i][j] = 0
			}
		}
	}
}

func processCell(board [][]int, row, col int) {
	score := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			}
			if i >= 0 && i < len(board) && j >= 0 && j < len(board[i]) && board[i][j]%2 == 1 {
				score++
			}
		}
	}

	val := board[row][col] % 2
	if (val == 0 && score == 3) || (val == 1 && (score < 2 || score > 3)) {
		board[row][col] += 2
	}
}

