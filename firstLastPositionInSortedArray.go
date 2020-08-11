package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	left := binarySearch(nums, target, true)

	if left == len(nums) || nums[left] != target {
		return result
	}

	right := binarySearch(nums, target, false) - 1

	return []int{left, right}

}

func binarySearch(nums []int, target int, low bool) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > target || nums[mid] == target && low {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
