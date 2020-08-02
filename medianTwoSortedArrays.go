// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.

package main

import "fmt"

func main() {
	a, b := []int{1, 3, 5, 7, 9}, []int{2, 4, 6}

	fmt.Println(findMedianSortedArrays(a, b))
}

// Time Complexity: O(min(m, n))
// Space Complexity: O(min(m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := len(nums1)-1, len(nums2)-1
	if i < 0 && j < 0 {
		return 0
	}

	var mergedArray []int
	if i < 0 {
		mergedArray = nums2
	} else if j < 0 {
		mergedArray = nums1
	} else if nums1[0] > nums2[0] {
		mergedArray = append(nums2, nums1...)
	} else {
		mergedArray = append(nums1, nums2...)
	}

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			mergedArray[i+j+1] = nums1[i]
			i--
		} else {
			mergedArray[i+j+1] = nums2[j]
			j--
		}
	}

	if len(mergedArray)%2 == 0 {
		return float64(mergedArray[len(mergedArray)/2]+mergedArray[len(mergedArray)/2-1]) / 2.0
	} else {
		return float64(mergedArray[len(mergedArray)/2])
	}
}

// Time complexity: O(log(min(m,n)))
// Space Complexity: O(1)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}

	min, max := 0, m

	for min <= max {
		i := (min + max) / 2
		j := (m+n+1)/2 - i

		if i < max && nums2[j-1] > nums1[i] { // i is too small
			min = i + 1
		} else if i > min && nums1[i-1] > nums2[j] { // i is too big
			max = i - 1
		} else { // i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = findMax(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0

			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = findMin(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}

	return 0.0
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
