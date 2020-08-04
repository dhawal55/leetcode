package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	resultMap := make(map[string]bool)
	tempResult := generateParenthesis(n - 1)

	current := ""
	for i := 0; i < len(tempResult); i++ {
		current = tempResult[i]
		for j := 0; j < len(current); j++ {
			temp := current[0:j+1] + "()" + current[j+1:]
			resultMap[temp] = true
		}
	}

	var result []string
	for k, _ := range resultMap {
		result = append(result, k)
	}

	return result
}
