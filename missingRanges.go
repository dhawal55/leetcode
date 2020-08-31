package main

import (
	"fmt"
	"math"
)

/*
Given a sorted integer array nums, where the range of elements are in the inclusive range [lower, upper], return its missing ranges.
Input: nums = [0, 1, 3, 50, 75], lower = 0 and upper = 99,
Output: ["2", "4->49", "51->74", "76->99"]
*/
func main() {
	nums := []int{0, 1, 3, 50, 75}
	lower, upper := 0, 99

	fmt.Println(findMissingRanges(nums, lower, upper))

	nums = []int{0, 0, 0, 1, 3, 50, 75}
	fmt.Println(findMissingRanges(nums, lower, upper))

	nums = []int{0, 1, 3, 50, 75, math.MaxInt64}
	lower, upper = 0, math.MaxInt64
	fmt.Println(findMissingRanges(nums, lower, upper))
}

func findMissingRanges(nums []int, lower, upper int) []string {
	var result []string
	if lower >= upper {
		return result
	}

	l, r := lower, upper
	for i := 0; i < len(nums); i++ {
		r = nums[i] - 1

		if l == r {
			result = append(result, fmt.Sprintf("%v", l))
		} else if l < r {
			result = append(result, fmt.Sprintf("%v-%v", l, r))
		}

		// Protect from int overflow
		if nums[i] == math.MaxInt64 {
			return result
		}
		l = nums[i] + 1
	}

	if l == upper {
		result = append(result, fmt.Sprintf("%v", l))
	} else if l < upper {
		result = append(result, fmt.Sprintf("%v-%v", l, upper))
	}

	return result
}
