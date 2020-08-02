// Given a string, find the length of the longest substring without repeating characters.
// Example 1:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
package main

import "fmt"

func main() {
	input := "p%wwk%ew"

	fmt.Println(lengthOfLongestSubstring(input))
}

// Sliding window
// Time complexity : O(2n) = O(n).
// Space complexity : O(min(m,n)).
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]bool)

	var left, right int
	var max int
	for right < len(s) {
		if _, ok := charMap[s[right]]; ok {
			if right-left > max {
				max = right - left
			}

			delete(charMap, s[left])
			left++
		} else {
			charMap[s[right]] = true
			right++
		}
	}

	if right-left > max {
		max = right - left
	}

	return max
}
