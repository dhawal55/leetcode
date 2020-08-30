package main

func main() {

}

// Time Complexity: O(m x n), where m is len of wordDict and n is len of s
// Space Complexity: O(n)
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{})
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if dp[i] {
				if i+len(word) <= len(s) && s[i:i+len(word)] == word {
					dp[i+len(word)] = true
				}
			}
		}
	}

	return dp[len(s)]
}

// Time Complexity: O(n^2) - I think so
// Space Complexity: O(n!)
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{})
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = struct{}{}
	}

	memoize := make(map[string]bool)
	return wordBreakRecurse(s, wordMap, 0, memoize)
}

func wordBreakRecurse(s string, wordMap map[string]struct{}, index int, memoize map[string]bool) bool {
	if index == len(s) {
		return true
	}

	if r, ok := memoize[s[index:]]; ok {
		return r
	}

	for i := index; i < len(s); i++ {
		str := s[index : i+1]

		if _, ok := wordMap[str]; ok {
			memoize[s[i+1:]] = wordBreakRecurse(s, wordMap, i+1, memoize)
			if memoize[s[i+1:]] {
				return true
			}
		}
	}

	memoize[s] = false
	return false
}
