package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(multiply("456", "123"))
}

func multiply(num1 string, num2 string) string {
	resultArr := make([]int8, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		a1 := int8(num1[i] - '0')

		for j := len(num2) - 1; j >= 0; j-- {
			a2 := int8(num2[j] - '0')

			sum := (a1 * a2) + resultArr[i+j+1]
			resultArr[i+j+1] = sum % 10
			resultArr[i+j] += sum / 10
		}
	}

	var result strings.Builder
	result.Grow(len(resultArr))
	index := 0
	for index < len(resultArr) && resultArr[index] == 0 {
		index++
	}

	for index < len(resultArr) {
		fmt.Fprintf(&result, "%v", resultArr[index])
		index++
	}

	if result.Len() == 0 {
		result.WriteString("0")
	}

	return result.String()
}
