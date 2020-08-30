package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return 0
	}

	var totalDiff, currDiff, start int
	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		totalDiff += diff
		currDiff += diff

		if currDiff < 0 {
			start = len(gas)
			currDiff = 0
		} else if i < start {
			start = i
		}
	}

	if start < len(gas) && totalDiff >= 0 {
		return start
	}

	return -1
}
