package main

import "fmt"

func main() {
	input := []int{1, 2, 3}
	fmt.Println(permute(input))
}

func permute(nums []int) [][]int {
	usedMap := make(map[int]struct{})
	return generate(nums, nil, nil, usedMap)
}

// Time Complexity: O(n ^ 2)
func generate(nums []int, prefix []int, result [][]int, usedMap map[int]struct{}) [][]int {
	if len(prefix) == len(nums) {
		return append(result, prefix)
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := usedMap[nums[i]]; ok {
			continue
		}
		usedMap[nums[i]] = struct{}{}
		result = generate(nums, append(prefix, nums[i]), result, usedMap)
		delete(usedMap, nums[i])
	}

	return result
}
