package main

import "fmt"

func main() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	wordArr := []byte(word)
	if len(board) == 0 || len(wordArr) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if match(board, wordArr, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func match(board [][]byte, word []byte, x, y, index int) bool {
	if index >= len(word) {
		return true
	}

	if x < 0 || x >= len(board) {
		return false
	}

	if y < 0 || y >= len(board[0]) {
		return false
	}

	if board[x][y] != word[index] {
		return false
	}

	temp := board[x][y]
	board[x][y] = ' '

	index++

	if match(board, word, x-1, y, index) || match(board, word, x+1, y, index) || match(board, word, x, y-1, index) || match(board, word, x, y+1, index) {
		return true
	}

	board[x][y] = temp
	return false
}
