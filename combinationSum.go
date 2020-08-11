package main

import (
	"fmt"
)

func main() {
	candidates := []int{7, 3, 2}
	target := 18

	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	return buildCombinationSum(candidates, target, nil)
}

func buildCombinationSum(candidates []int, target int, combination []int) [][]int {
	var result [][]int

	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] > target {
			continue
		}

		newCombination := make([]int, len(combination)+1)
		copy(newCombination, combination)
		newCombination[len(combination)] = candidates[i]

		if candidates[i] == target {
			result = append(result, newCombination)
		} else {
			result = append(result, buildCombinationSum(candidates[0:i+1], target-candidates[i], newCombination)...)
		}
	}

	return result
}
