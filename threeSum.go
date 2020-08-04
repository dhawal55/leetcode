package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))
}

// Time Complexity: O(nlogn + n^2) = O(n^2) (sort operation doesn't change overall time complexity)
// Space Complexity: O(n) to O(logn) for sort algo
// Two Pointer for 2Sum
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		sum := 0 - nums[i]
		left, right := i+1, len(nums)-1

		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			currentSum := nums[left] + nums[right]
			if currentSum == sum {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if currentSum < sum {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
