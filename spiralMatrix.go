package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{0, 10},
		[]int{1, 11},
		[]int{2, 12},
		[]int{3, 13},
		[]int{4, 14},
		[]int{5, 15},
	}

	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	r1, r2 := 0, len(matrix)-1
	c1, c2 := 0, len(matrix[0])-1

	var result []int
	for r1 <= r2 && c1 <= c2 {
		for c := c1; c <= c2; c++ {
			result = append(result, matrix[r1][c])
		}

		for r := r1 + 1; r <= r2; r++ {
			result = append(result, matrix[r][c2])
		}

		if r1 < r2 && c1 < c2 {
			for c := c2 - 1; c > c1; c-- {
				result = append(result, matrix[r2][c])
			}

			for r := r2; r > r1; r-- {
				result = append(result, matrix[r][c1])
			}
		}

		r1++
		r2--
		c1++
		c2--
	}

	return result
}
