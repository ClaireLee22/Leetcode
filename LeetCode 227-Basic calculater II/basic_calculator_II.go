package main

import "fmt"

func calculate(s string) int {
	finalResult, tempResult, operand := 0, 0, 0
	operator := '+'

	for idx, char := range s {
		if '0' <= char && char <= '9' {
			operand = operand*10 + int(char-'0')
		}

		if idx == len(s)-1 ||
			char == '+' ||
			char == '-' ||
			char == '*' ||
			char == '/' {
			if operator == '-' || operator == '+' {
				finalResult += tempResult
				tempResult = operand
				if operator == '-' {
					tempResult = -operand
				}
			} else if operator == '*' {
				tempResult *= operand
			} else if operator == '/' {
				tempResult /= operand
			}

			operand = 0
			operator = char
		}
	}

	finalResult += tempResult

	return finalResult
}

func main() {
	fmt.Println("result:", calculate("10 + 2*3 - 6/3"))
}

/*
output:
result: 14
*/
