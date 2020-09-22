package main() 

func main() {
	coins := []int{186,419,83,408}
	amount := 6249

	fmt.Println(coinChange(amount, coins))
}

// DP - bottom-up tabulation, O(N x M), where N is amount & M is # of coins
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 0; i <= amount; i++ {
			dp[i] = amount+1
	}
	
	dp[0] = 0
	for i := 1; i <= amount; i++ {
			for _, coin := range coins {
					if coin <= i {
							dp[i] = int(math.Min(float64(1 + dp[i-coin]), float64(dp[i])))
					}
			}
	}
	
	if dp[amount] == amount+1 {
			return -1
	}
	return dp[amount]
}

// DP - top-dwon memoization, O(N x M), where N is amount & M is # of coins
func coinChangeTopDown(coins []int, amount int) int {
	dp := make([]int, amount+1)
	return coinChangeRecurse(coins, amount, dp)
}

func coinChangeRecurse(coins []int, amount int, dp []int) int {
	if amount < 0 {
			return -1
	}
	
	if amount == 0 {
			return 0
	}

	if dp[amount] != 0 {
			return dp[amount]
	}
	
	minTotal := amount + 1
	for _, coin := range coins {
			total := coinChangeRecurse(coins, amount - coin, dp)
					
			if total >= 0 && total < minTotal {
					minTotal = 1 + total
			}
	}
	
	if minTotal > amount {
			minTotal = -1
	}
	dp[amount] = minTotal
	
	return minTotal
}