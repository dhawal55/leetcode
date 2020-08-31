package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
}

func findPeakElement(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
