package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
	fmt.Println(groupAnagramsNormalSort(input))
}

// TimeComplexity: O(nm), n is len of input & k is max len of a string
// SpaceComplexity: O(nm)
func groupAnagrams(strs []string) [][]string {
	sortedMap := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		// Count sorting
		count := [26]int8{}
		for j := 0; j < len(strs[i]); j++ {
			count[strs[i][j]-'a']++
		}

		var sb strings.Builder
		for j := 0; j < len(count); j++ {
			fmt.Fprintf(&sb, "%d", count[j])
		}

		sortedMap[sb.String()] = append(sortedMap[sb.String()], strs[i])
	}

	var result [][]string
	for _, arr := range sortedMap {
		result = append(result, arr)
	}

	return result
}

// TimeComplexity: O(nmlogm)
// SpaceComplexity: O(nm)
func groupAnagrams(strs []string) [][]string {
	sortedMap := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		strArr := []byte(strs[i])
		sort.Slice(strArr, func(i, j int) bool { return strArr[i] < strArr[j] })

		str := string(strArr)
		sortedMap[str] = append(sortedMap[str], strs[i])
	}

	var result [][]string
	for _, arr := range sortedMap {
		result = append(result, arr)
	}

	return result
}
