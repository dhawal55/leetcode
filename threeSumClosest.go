package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1

	fmt.Println(threeSumClosest(nums, target))
}

// Time Complexity: O(nlogn + n^2)= O(n^2)
// Space Complexity: O(n) to O(logn) for sort algo
func threeSumClosest(nums []int, target int) int {
	var diff, output int
	minDiff := math.MaxInt32

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right] + nums[i]

			if sum == target {
				return sum
			} else if sum > target {
				right--
			} else {
				left++
			}

			if sum > target {
				diff = sum - target
			} else {
				diff = target - sum
			}

			if diff < minDiff {
				minDiff = diff
				output = sum
			}
		}

	}

	return output
}
