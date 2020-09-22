package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 1, 1, 2, 2, 1}
	wiggleSort(nums)
	fmt.Println(nums)
}

func wiggleSort(nums []int) {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	clone := make([]int, len(nums))
	copy(clone, nums)

	mid := len(clone) / 2
	j := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = clone[mid]
			mid++
		} else {
			nums[i] = clone[j]
			j++
		}
	}
}
