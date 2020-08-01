package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

// func convert(s string, numRows int) string {
//     rows := make([]string, numRows)

//     if numRows == 1 {
//         return s
//     }

//     goingDown := false
//     curRow := 0
//     inputRune := []rune(s)

//     for _, r := range inputRune {
//         if curRow == 0 || curRow == numRows -1 {
//             goingDown = !goingDown
//         }

//         rows[curRow] = rows[curRow] + string(r)
//         if goingDown {
//             curRow++
//         } else {
//             curRow--
//         }
//     }

//     var output string
//     for _, v := range rows {
//         output = output + v
//     }

//     return output
// }

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
