package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}

	fmt.Println(searchMatrix(matrix, 15))
}
func searchMatrix(m [][]int, target int) bool {
	for i, j := len(m)-1, 0; i >= 0 && j < len(m[0]); {
		if m[i][j] > target {
			i--
		} else if m[i][j] < target {
			j++
		} else {
			return true
		}
	}

	return false
}
