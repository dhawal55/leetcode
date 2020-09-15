//Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] find the minimum number of conference rooms required.
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		[]int{13, 16},
		[]int{8, 14},
		[]int{9, 11},
		[]int{15, 17},
		[]int{12, 16},
		[]int{8, 10},
	}

	fmt.Println(minMeetingRooms(intervals))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= (*h)[0] {
			heap.Pop(h)
		}

		heap.Push(h, intervals[i][1])
	}

	return len(*h)
}
