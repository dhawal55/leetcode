// Given an unsorted array of integers, find the length of longest increasing subsequence.
// Example:
// Input: [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
// Note:

// There may be more than one LIS combination, it is only necessary for you to return the length.
// Your algorithm should run in O(n2) complexity.
// Follow up: Could you improve it to O(n log n) time complexity?
package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(longestIncreasingSubsequence(nums))
}

func longestIncreasingSubsequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSubsequence := make([]int, len(nums))
	maxAns := 0

	for i := 1; i < len(nums); i++ {
		maxVal := 0

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxVal = max(maxVal, maxSubsequence[j])
			}
		}

		maxSubsequence[i] = maxVal + 1
		maxAns = max(maxAns, maxSubsequence[i])
	}

	return maxAns
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}

	return val2
}
