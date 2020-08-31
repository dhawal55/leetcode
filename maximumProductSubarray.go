package main

import "fmt"

func main() {
	nums := []int{-4, -3, -2}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	min := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		oldMax := max
		if max*nums[i] > min*nums[i] {
			max = findMax(nums[i], max*nums[i])
			min = findMin(nums[i], min*nums[i])
		} else {
			max = findMax(nums[i], min*nums[i])
			min = findMin(nums[i], oldMax*nums[i])
		}

		result = findMax(result, findMax(max, min))
	}

	return result
}

func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
