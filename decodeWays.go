package main

import "fmt"

func main() {
	fmt.Println(numDecodings("120"))
	fmt.Println(numDecodings("1200"))
	fmt.Println(numDecodings("17"))
	fmt.Println(numDecodings("12121"))
	fmt.Println(numDecodings("12321"))
}

func numDecodings(s string) int {
	byteSlice := []byte(s)

	var count [3]int
	if byteSlice[len(byteSlice)-1] != '0' {
		count[0] = 1
		count[1] = 1
	}

	for i := len(byteSlice) - 2; i >= 0; i-- {
		if byteSlice[i] == '0' {
			count[0] = 0
		} else {
			count[0] = count[1]
			if int(byteSlice[i]-'0')*10+int(byteSlice[i+1]-'0') <= 26 {
				if i == len(byteSlice)-2 {
					count[0] += 1
				} else {
					count[0] += count[2]
				}
			}
		}

		count[2] = count[1]
		count[1] = count[0]
	}

	return count[0]
}
