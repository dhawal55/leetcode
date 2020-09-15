package main

import (
	"fmt"
	"math"
)

func main() {
	s := "    -42"
	fmt.Println(myAtoi(s))
}

// Time Complexity; O(n), where n is len of string
func myAtoi(str string) int {
	// Trim white space
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}

	var negative bool
	// Check for negative/positive sign
	if i < len(str) && (str[i] == '-' || str[i] == '+') {
		negative = (str[i] == '-')
		i++
	}

	var result int
	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}

		// Use MaxInt64 on a 64-bit OS
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && str[i]-'0' > math.MaxInt32%10) {
			if negative {
				return math.MinInt32
			}
			return math.MaxInt32
		}

		result = (result * 10) + int(str[i]-'0')
	}

	if negative {
		return 0 - result
	}

	return result
}
