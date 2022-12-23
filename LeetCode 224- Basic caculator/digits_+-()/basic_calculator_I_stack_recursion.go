package main

import "fmt"

func calculate(s string) int {
	array := make([]rune, len(s))
	for idx, char := range s {
		array[idx] = char
	}
	return calculateHelper(&array)
}

func calculateHelper(array *[]rune) int {
	operand := 0
	operator := 1 // +
	stack := []int{}

	for len(*array) > 0 {
		char := (*array)[0]
		*array = (*array)[1:]

		if '0' <= char && char <= '9' {
			operand = operand*10 + int(char-'0')
		} else if char == '+' || char == '-' {
			stack = append(stack, operand*operator)
			operand = 0
			operator = 1
			if char == '-' {
				operator = -1
			}
		} else if char == '(' {
			operand = calculateHelper(array)
		} else if char == ')' {
			break
		}

	}
	return sum(stack) + operand*operator
}

func sum(array []int) int {
	total := 0
	for _, num := range array {
		total += num
	}
	return total
}

func main() {
	s := "10+(2-3+5)-6"
	fmt.Println("result:", calculate(s))
}

/*
output:
result: 8
*/
