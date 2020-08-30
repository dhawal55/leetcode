package main

import "fmt"

func main() {
	fmt.Println(partition("cbbbcc"))
}

func partition(s string) [][]string {
	var result [][]string
	if len(s) == 0 {
		return result
	}

	return partitionRecurse(0, s, nil, result)
}

func partitionRecurse(index int, s string, current []string, result [][]string) [][]string {
	if index == len(s) {
		clone := make([]string, len(current))
		copy(clone, current)
		result = append(result, clone)
	}

	for i := index; i < len(s); i++ {
		if isPalindrome(s[index : i+1]) {
			current = append(current, string(s[index:i+1]))
			result = partitionRecurse(i+1, s, current, result)
			current = current[:len(current)-1]
		}
	}

	return result
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}
