package main

import "fmt"

func main() {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}

	solve(board)
	fmt.Println(board)
}

// Time Complexity: O(MN), M is # of rows and N is # of cols
// Space Complexity: O(1)
func solve(board [][]byte) {
	if len(board) < 3 {
		return
	}

	if len(board[0]) < 3 {
		return
	}

	rows, cols := len(board), len(board[0])

	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}

		if board[i][cols-1] == 'O' {
			dfs(board, i, cols-1)
		}
	}

	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}

		if board[rows-1][j] == 'O' {
			dfs(board, rows-1, j)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'E' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}

	board[i][j] = 'E'

	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}
