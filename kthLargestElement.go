package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{3, 6, 2, 8, 9, 1, 6}
	k := 2

	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}

	quickSort(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

func quickSort(nums []int, low, high, k int) {
	if low < high {
		pivot := partition(nums, low, high)
		if pivot == k-1 {
			return
		} else if pivot > k-1 {
			quickSort(nums, low, pivot-1, k)
		} else {
			quickSort(nums, pivot+1, high, k)
		}
	}
}

func partition(nums []int, low, high int) int {
	rand.Seed(time.Now().UnixNano())
	pivot := low + rand.Int()%(high-low)

	nums[pivot], nums[high] = nums[high], nums[pivot]
	for i := low; i < high; i++ {
		if nums[i] > nums[high] {
			nums[i], nums[low] = nums[low], nums[i]
			low++
		}
	}

	nums[low], nums[high] = nums[high], nums[low]
	return low
}
