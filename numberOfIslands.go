package main

import "fmt"

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}

			checkSurroundedByWater(grid, i, j)
			count++
		}
	}

	return count
}

func checkSurroundedByWater(grid [][]byte, i, j int) {
	grid[i][j] = '0'

	// check top
	if i > 0 && grid[i-1][j] == '1' {
		checkSurroundedByWater(grid, i-1, j)
	}

	// check bottom
	if i < len(grid)-1 && grid[i+1][j] == '1' {
		checkSurroundedByWater(grid, i+1, j)
	}

	// check left
	if j > 0 && grid[i][j-1] == '1' {
		checkSurroundedByWater(grid, i, j-1)
	}

	// check right
	if j < len(grid[i])-1 && grid[i][j+1] == '1' {
		checkSurroundedByWater(grid, i, j+1)
	}
}
