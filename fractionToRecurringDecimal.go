package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(fractionToDecimal(4, 333))
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 || denominator == 0 {
		return "0"
	}

	var negative bool
	if numerator < 0 {
		numerator = 0 - numerator
		negative = !negative
	}

	if denominator < 0 {
		denominator = 0 - denominator
		negative = !negative
	}

	var result strings.Builder
	if negative {
		fmt.Fprintf(&result, "-")
	}

	var count int
	count, numerator = divide(numerator, denominator)

	if numerator == 0 {
		fmt.Fprintf(&result, "%v", count)
		return result.String()
	}

	fmt.Fprintf(&result, "%v.", count)
	index := result.Len()

	dp := make(map[int]int)
	for numerator > 0 {
		if index, ok := dp[numerator]; ok {
			s := result.String()
			return fmt.Sprintf("%s(%s)", s[0:index], s[index:])
		}

		dp[numerator] = index
		// Multiply by 10
		numerator = numerator<<1 + numerator<<3
		count, numerator = divide(numerator, denominator)
		fmt.Fprintf(&result, "%v", count)
		index++
	}

	return result.String()
}

func divide(numerator int, denominator int) (int, int) {
	var result int

	for numerator >= denominator {
		d := denominator
		count := 1

		for numerator > d<<1 {
			d = d << 1
			count = count << 1
		}

		numerator = numerator - d
		result = result + count
	}

	return result, numerator
}

func fractionToDecimalMath(numerator int, denominator int) string {
	if numerator == 0 || denominator == 0 {
		return "0"
	}

	var result strings.Builder
	if (numerator < 0) != (denominator < 0) {
		fmt.Fprintf(&result, "-")
	}

	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))

	fmt.Fprintf(&result, "%v", numerator/denominator)
	numerator = numerator % denominator
	if numerator == 0 {
		return result.String()
	}

	fmt.Fprintf(&result, ".")
	dp := make(map[int]int)

	index := result.Len()
	for numerator > 0 {
		if index, ok := dp[numerator]; ok {
			s := result.String()
			return fmt.Sprintf("%s(%s)", s[0:index], s[index:])
		}

		dp[numerator] = index
		numerator = numerator * 10
		fmt.Fprintf(&result, "%v", numerator/denominator)
		numerator = numerator % denominator
		index++
	}

	return result.String()
}
