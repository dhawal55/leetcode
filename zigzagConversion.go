package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

// Time Complexity: O(n), n is len of string. Each index is visited once.
// Space Complexity: O(n) or O(1) if return string is not considered extra space
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	r := []rune(s)
	cycleLen := 2*numRows - 2
	var output []rune
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(r); j += cycleLen {
			output = append(output, r[j+i])

			if i != 0 && i != numRows-1 && j+cycleLen-i < len(s) {
				output = append(output, r[j+cycleLen-i])
			}
		}
	}

	return string(output)
}
