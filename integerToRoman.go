package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	romanNums := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M', 'M', 'M'}

	if num > 3999 {
		return "ERROR"
	}

	var result []byte
	count := 0

	for num > 0 {
		n := num % 10
		i := n % 5
		j := n / 5
		temp := []byte{}

		if i == 4 {
			temp = []byte{romanNums[count], romanNums[count+j+1]}
		} else {
			if j > 0 {
				temp = []byte{romanNums[count+1]}
			}

			for k := 1; k <= i; k++ {
				temp = append(temp, romanNums[count])
			}
		}

		count += 2
		result = append(temp, result...)
		num = num / 10
	}

	return string(result)
}

func intToRoma1(num int) string {
	romanNums := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	if num > 3999 {
		return "ERROR"
	}

	var result string

	for i := 0; i < len(nums); i++ {
		t := num / nums[i]
		num = num % nums[i]

		for j := 0; j < t; j++ {
			result = result + romanNums[i]
		}

		if num == 0 {
			break
		}
	}

	return string(result)
}
