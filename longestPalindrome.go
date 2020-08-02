// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
// Example: Input: "babad", Output: "bab", Input: "cbbd", Output: "bb"

package main

import (
	"fmt"
)

func main() {
	s := "eaabbaaed"
	fmt.Println(longestPalindrome(s))
}

// Expand around center
// Time Complexity:  O(n^2)
// Space Complexity: O(1)

// TODO: O(n) Non-trivial Machester algo
func longestPalindrome(input string) string {
	if len(input) == 0 {
		return ""
	}

	inputRune := []rune(input)
	var longestLow, longestHigh int

	for i := 0; i < len(inputRune); i++ {
		low, high := i-1, i+1

		for low >= 0 && inputRune[low] == inputRune[i] {
			low--
		}

		for low >= 0 && high < len(inputRune) && inputRune[low] == inputRune[high] {
			low--
			high++
		}

		if ((high - 1) - (low + 1)) > (longestHigh - longestLow) {
			longestLow = low + 1
			longestHigh = high - 1
		}
	}

	return string(inputRune[longestLow : longestHigh+1])
}
