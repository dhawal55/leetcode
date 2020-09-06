package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}

// Time Complexity: O(nlogn)
// Space Complexity: O(n) - we must allocate O(n) space for the final return string
func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	sort.Slice(nums, func(i, j int) bool {
		return fmt.Sprintf("%v%v", nums[i], nums[j]) > fmt.Sprintf("%v%v", nums[j], nums[i])
	})
	if nums[0] == 0 {
		return "0"
	}

	s := strings.Builder{}
	for i := 0; i < len(nums); i++ {
		fmt.Fprintf(&s, "%v", nums[i])
	}

	return s.String()
}

// Custom sort if you can't use sort lib
func sortSlice(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)

		sortSlice(nums, left, pivot-1)
		sortSlice(nums, pivot+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := right

	for i := left; i < right; i++ {
		if greaterThan(nums[i], nums[pivot]) {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[pivot] = nums[pivot], nums[left]
	return left
}

func greaterThan(i, j int) bool {
	str1, str2 := fmt.Sprintf("%v%v", i, j), fmt.Sprintf("%v%v", j, i)
	return str1 > str2
}
