package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	input := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(input))
}

func evalRPN(tokens []string) int {
	var queue []int

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			queue[len(queue)-2] = queue[len(queue)-2] + queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		case "-":
			queue[len(queue)-2] = queue[len(queue)-2] - queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		case "*":
			queue[len(queue)-2] = queue[len(queue)-2] * queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		case "/":
			queue[len(queue)-2] = queue[len(queue)-2] / queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		default:
			num, err := strconv.Atoi(tokens[i])
			if err != nil {
				log.Fatal("Invalid operator", tokens[i])
			}
			queue = append(queue, num)
		}
	}

	return queue[0]
}
