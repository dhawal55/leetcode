package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4, 5}))
}

// Time complexity: O(N x 2^N) to generate all subsets and then copy them into output list.
// Space complexity: O(N×2^N) to keep all the subsets of length N, since each of N elements could be present or absent.
func subsets(nums []int) [][]int {
	result := [][]int{
		[]int{},
	}

	for i := 0; i < len(nums); i++ {
		count := len(result)
		for j := 0; j < count; j++ {
			r := len(result[j])
			dup := make([]int, r+1)
			copy(dup, result[j])
			dup[r] = nums[i]
			result = append(result, dup)
		}
	}

	return result
}
