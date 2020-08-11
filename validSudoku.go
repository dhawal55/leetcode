package main

import "fmt"

func main() {
	sudoku := [][]byte{
		[]byte{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		[]byte{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		[]byte{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		[]byte{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		[]byte{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		[]byte{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}

	fmt.Println(isValidSudoku(sudoku))
}

type Position struct {
	Row, Col byte
}

func isValidSudoku(board [][]byte) bool {
	positionMap := make(map[byte][]Position)
	size := byte(9)

	for i := byte(0); i < size; i++ {
		for j := byte(0); j < size; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			positions := positionMap[num]
			subboard := (i / 3 * 3) + (j / 3)

			for _, p := range positions {
				if p.Row == i || p.Col == j || subboard == (p.Row/3*3)+(p.Col/3) {
					return false
				}
			}

			positionMap[num] = append(positions, Position{Row: i, Col: j})
		}
	}

	return true
}
