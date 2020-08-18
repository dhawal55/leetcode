package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 4))
}

// Time Complexity: O(m.n)
// Space Complexity: O(m.n)
func uniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}
