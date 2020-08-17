package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18},
	}

	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	var result [][]int
	min, max := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= min && intervals[i][0] <= max {
			if intervals[i][1] > max {
				max = intervals[i][1]
			}
		} else {
			result = append(result, []int{min, max})
			min = intervals[i][0]
			max = intervals[i][1]
		}
	}

	result = append(result, []int{min, max})
	return result
}
