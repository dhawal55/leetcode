package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(nums, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return buildCombinationSum2(candidates, target, nil)
}

func buildCombinationSum2(candidates []int, target int, combination []int) [][]int {
	var result [][]int

	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}

		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}

		newCombination := make([]int, len(combination)+1)
		copy(newCombination, combination)
		newCombination[len(combination)] = candidates[i]

		if target-candidates[i] == 0 {
			result = append(result, newCombination)
		} else {
			result = append(result, buildCombinationSum2(candidates[i+1:], target-candidates[i], newCombination)...)
		}
	}

	return result
}
