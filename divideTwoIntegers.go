package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(math.MaxInt32, 2))
}

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	absDividend, absDivisor := int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	var ret int
	for absDividend >= absDivisor {
		acc, cnt := absDivisor, 1
		for absDividend-(acc<<1) > 0 {
			acc, cnt = acc<<1, cnt<<1
		}
		absDividend, ret = absDividend-acc, ret+cnt
	}

	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		return -ret
	}
	return ret
}
