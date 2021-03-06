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
// Space complexity : O(n)
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]struct{})

	var left, right, max int
	for right < len(s) {
		if _, ok := charMap[s[right]]; ok {
			if right-left > max {
				max = right - left
			}

			delete(charMap, s[left])
			left++
		} else {
			charMap[s[right]] = struct{}{}
			right++
		}
	}

	if right-left > max {
		max = right - left
	}

	return max
}

// Time complexity : O(n). Index right will iterate n times.
// Space complexity: O(min(m,n)). We need O(k) space for checking a substring has no duplicate characters, where k is the size of the charArr. The size of the array is upper bounded by the size of the string n and the size of the charset/alphabet m.
func lengthOfLongestSubstring(s string) int {
	// [26]int for letters 'a' - 'z' or 'A' - 'Z'
	// [128]int for ASCII
	// [256]int fot Extended ASCII
	var charArr [128]int

	var left, right, max int
	for right < len(s) {
		if charArr[s[right]] > left {
			left = charArr[s[right]]
		}

		if right-left+1 > max {
			max = right - left + 1
		}
		charArr[s[right]] = right + 1
		right++
	}

	return max
}
