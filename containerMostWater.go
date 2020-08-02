package main

import "fmt"

func main() {
	height := []int{1, 4, 6, 3, 5, 7, 9, 5, 3, 1}
	fmt.Println(maxArea(height))
}

// Time Complexity: O(n)
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	capacity := 0

	for left < right {
		i, j := left, right

		if height[left] < height[right] {
			capacity = height[left] * (right - left)
			left++
		} else {
			capacity = height[right] * (right - left)
			right--
		}

		if capacity > max {
			max = capacity
		}
	}

	return max
}
