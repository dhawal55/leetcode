package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{}
	target := 0
	fmt.Println(fourSum(nums, target))
}

// Time Complexity: O(n^3)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 0, 4)
}

// Time Complexity: O(n^k−1)
// Space Complexity: O(n). We need O(k) space for the recursion. k can be the same as n in the worst case for the generalized algorithm.
func kSum(nums []int, target, start, k int) [][]int {
	var result [][]int

	if start == len(nums) || nums[start]*k > target || nums[len(nums)-1]*4 < target {
		return result
	}

	if k == 2 {
		return twoSum(nums, target, start)
	}

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		temp := kSum(nums, target-nums[i], i+1, k-1)
		for j := 0; j < len(temp); j++ {
			result = append(result, append(temp[j], nums[i]))
		}
	}

	return result
}

// Sliding pattern
func twoSum(nums []int, target, start int) [][]int {
	var result [][]int

	left, right := start, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]

		if (left > start && nums[left] == nums[left-1]) || sum < target {
			left++
		} else if (right < len(nums)-1 && nums[right] == nums[right+1]) || sum > target {
			right--
		} else {
			result = append(result, []int{nums[left], nums[right]})
			left++
			right--
		}
	}

	return result
}
