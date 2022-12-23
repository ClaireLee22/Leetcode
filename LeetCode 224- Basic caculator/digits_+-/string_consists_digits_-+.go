package main

import "fmt"

func calculate(s string) int {
	result := 0
	operand := 0
	operator := 1

	for _, char := range s {
		if '0' <= char && char <= '9' {
			operand = operand*10 + int(char-'0')
		} else if char == '+' || char == '-' {
			result += operand * operator
			operand = 0
			operator = 1
			// save the operator we currently met for the next evaluation
			if char == '-' {
				operator = -1
			}
		}
	}
	result += operand * operator

	return result
}

func main() {
	s := "10+4-6"
	fmt.Println("result:", calculate(s))
}

/*
output:
result: 8
*/
