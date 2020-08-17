package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

// Time Complexity = O(n)
// Space Complexity = O(1)
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	lastPos := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func canJumpSlow(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	canVisit := make([]bool, len(nums))
	canVisit[0] = true

	for i := 0; i < len(nums)-1 && canVisit[i] == true; i++ {
		for j := i + 1; j <= nums[i]+i; j++ {
			if j >= len(nums)-1 {
				return true
			}
			canVisit[j] = true
		}
	}

	return canVisit[len(nums)-1]
}
