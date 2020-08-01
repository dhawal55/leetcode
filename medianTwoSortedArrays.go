// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.

package main

import "fmt"

func main() {
	a, b := []int{1, 3, 5, 7, 9}, []int{2, 4, 6}

	fmt.Println(findMedianSortedArrays(a, b))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return (median(nums1) + median(nums2)) / 2.0
}

func median(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	var median float64
	if len(nums)%2 == 0 {
		median = float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2.0
	} else {
		median = float64(nums[len(nums)/2])
	}

	return median
}
