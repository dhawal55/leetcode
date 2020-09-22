package main

import "fmt"

func main() {
	fmt.Println(numSquares(12))
}

// Subproblem: The least number of perfect square numbers which sum to n is 1 + the least number of perfect square numbers which sum to (n - square number)
// We know the answer for n=0 and n=1, start from 2 and build it all the way to n
// TimeComplexity: O(n * sqrt(n))
func numSquares(n int) int {
	dp := make([]int, n+1)

	if n < 2 {
		return n
	}

	// Fill the array with n which is maximum number of squares numbers (1) that sum to n
	for i := 0; i < len(dp); i++ {
		dp[i] = n
	}

	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		// Loop through till the perfect square numbers is less than i
		for j := 1; j*j <= i; j++ {
			if 1+dp[i-j*j] < dp[i] {
				dp[i] = 1 + dp[i-j*j]
			}
		}
	}

	fmt.Println(dp)
	return dp[n]
}
