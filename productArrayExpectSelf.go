package main

import "fmt"

func main() {
	nums := []int{0, 5, 1}

	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	results := make([]int, len(nums))
	results[0] = 1

	// Compute product of all elements to the left
	for i := 1; i < len(nums); i++ {
		results[i] = results[i-1] * nums[i-1]
	}

	right := 1
	// Compute product of all elements to the right
	for i := len(nums) - 1; i >= 0; i-- {
		results[i] = results[i] * right
		right = right * nums[i]
	}

	return results
}

func productExceptSelf2(nums []int) []int {
	product := 1
	var zeroCount int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			product = product * nums[i]
		} else {
			zeroCount++
			if zeroCount > 2 {
				break
			}
		}
	}

	result := make([]int, len(nums))
	if zeroCount > 1 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result[i] = product
		} else if zeroCount > 0 {
			result[i] = 0
		} else {
			result[i] = product / nums[i]
		}
	}

	return result
}
