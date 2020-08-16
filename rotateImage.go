package main

import "fmt"

func main() {
	image := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	rotate(image)
	fmt.Println(image)
}

func rotate(matrix [][]int) {
	if len(matrix) < 2 {
		return
	}

	var current, next int
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			row, col := i, j
			current = matrix[row][col]

			for k := 0; k < 4; k++ {
				row, col = col, len(matrix)-row-1

				next = matrix[row][col]
				matrix[row][col] = current
				current = next
			}
		}
	}
}
