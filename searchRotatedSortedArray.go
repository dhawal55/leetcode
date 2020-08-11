package main

import "fmt"

func main() {
	// nums := []int{5, 1, 2, 3, 4}
	// target := 5

	// nums := []int{8, 9, 10, 1, 2, 3, 4, 5, 6}
	// target := 6

	nums := []int{1, 2, 3, 4, 5, 6, 0}
	target := 6
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// Left half of the array is sorted
		if nums[left] <= nums[mid] {
			// Usual binary search
			if target < nums[left] || target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // Right half of the array is sorted
			// Usual binary search
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
