package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

var numLetterMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	return letterCombinationRecursive(digits, "")
}

// Time complexity : O(3^N * 4^M) where N is the number of digits in the input that maps
//		to 3 letters (e.g. 2, 3, 4, 5, 6, 8) and M is the number of digits in the input
// 		that maps to 4 letters (e.g. 7, 9), and N+M is the total number digits in the input.
// Space complexity : O(3^N * 4^M) since one has to keep 3^N * 4^M solutions.
func letterCombinationRecursive(digits, prefix string) []string {
	var result []string
	if len(digits) == 0 {
		if len(prefix) > 0 {
			return []string{prefix}
		} else {
			return result
		}
	}

	if letters, ok := numLetterMap[digits[0]]; ok {
		for j := 0; j < len(letters); j++ {
			result = append(result, letterCombinationRecursive(digits[1:], prefix+string(letters[j]))...)
		}
	}

	return result
}
