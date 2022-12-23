package main

import "fmt"

func calculate(s string) int {
	result := 0
	operand := 0
	operator := 1 // +
	stack := []int{}

	for _, char := range s {
		if '0' <= char && char <= '9' {
			operand = operand*10 + int(char-'0')
		} else if char == '-' || char == '+' {
			result += operand * operator
			// reset operand
			operand = 0
			// update operator
			operator = 1
			if char == '-' {
				operator = -1
			}
		} else if char == '(' {
			stack = append(stack, result)
			stack = append(stack, operator)
			// reset result and operator
			result = 0
			operator = 1
		} else if char == ')' {
			result += operand * operator
			result *= stack[len(stack)-1]
			result += stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			// reset operand
			operand = 0
		}
	}

	result += operand * operator
	return result
}

func main() {
	s := "10+(2-3+5)-6"
	fmt.Println("result:", calculate(s))
}

/*
output:
result: 8
*/
