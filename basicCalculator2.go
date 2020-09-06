package main

import "fmt"

func main() {
	fmt.Println(calculate("1-1+1"))
	fmt.Println(calculate("   3  / 2 + 1    * 6  "))
}

func calculate(s string) int {
	var operandList []int
	prevSign := byte('+')
	var operand int

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			operand = operand*10 + int(s[i]-'0')
			continue
		}

		operandList = performOperation(operand, prevSign, operandList)
		prevSign = s[i]
		operand = 0
	}

	operandList = performOperation(operand, prevSign, operandList)
	var result int
	for i := 0; i < len(operandList); i++ {
		result += operandList[i]
	}

	return result
}

func calculateWithoutStack(s string) int {
	add, mul := 1, 1
	// op1 only used for add/sub
	op1, op2 := 0, 1

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			num := int(s[i] - '0')
			for i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
				i++
				num = num*10 + int(s[i]-'0')
			}

			if mul == 1 {
				op2 = op2 * num
			} else {
				op2 = op2 / num
			}
		} else if s[i] == '+' || s[i] == '-' {
			// Set op1 as op1 + op2 or op1 - op2 based on value of add var
			op1 += add * op2
			// Reset op2 and mul operator
			op2, mul = 1, 1

			add = 1
			if s[i] == '-' {
				add = -1
			}
		} else if s[i] == '*' || s[i] == '/' {
			mul = 1
			if s[i] == '/' {
				mul = -1
			}
		}
	}

	return op1 + add*op2
}

func performOperation(operand int, prevSign byte, operandList []int) []int {
	switch prevSign {
	case '+':
		operandList = append(operandList, operand)
	case '-':
		operandList = append(operandList, -operand)
	case '*':
		operandList[len(operandList)-1] = operandList[len(operandList)-1] * operand
	case '/':
		operandList[len(operandList)-1] = operandList[len(operandList)-1] / operand
	}

	return operandList
}
