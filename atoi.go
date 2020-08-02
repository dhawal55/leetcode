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
	var out int
	var negative bool

	i := 0
	for i < len(str) {
		if str[i] == ' ' {
			i++
		} else if str[i] == '-' || str[i] == '+' {
			negative = (str[i] == '-')
			i++
			break
		} else {
			break
		}
	}

	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}

		if (out*10)+int(str[i]-'0') > math.MaxInt32 {
			if negative {
				return math.MinInt32
			}
			return math.MaxInt32
		}

		out = (out * 10) + int(str[i]-'0')
	}

	if negative {
		return 0 - out
	}

	return out
}
