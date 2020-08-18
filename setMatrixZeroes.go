package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{1, 2, 3, 0},
		[]int{4, 5, 0, 6},
		[]int{7, 8, 9, 10},
	}

	setZeroes(matrix)
	fmt.Println(matrix)

	setZeroesMoreMemory(matrix)
	fmt.Println(matrix)
}

// Time Complexity: O(m.n)
// Space Complexity: O(1)
func setZeroes(matrix [][]int) {
	var firstColZero bool

	// Set 1st item in row/col to zero if the entire row/col needs to be set to zero
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
		}

		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// Set row/col to zero if either 1st item in row/col is zero
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Set 1st row to zero if needed
	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	// Set 1st col to zero if needed
	if firstColZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

// Time Complexity: O(m.n)
// Space Complexity: O(m+n)
func setZeroesMoreMemory(matrix [][]int) {
	var rows, cols []int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if contains(rows, i) || contains(cols, j) {
				matrix[i][j] = 0
			}
		}
	}
}

func contains(arr []int, val int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}

	return false
}
