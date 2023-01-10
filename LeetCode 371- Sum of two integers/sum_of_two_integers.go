package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func main() {
	a, b := 2, 3
	fmt.Println(getSum(a, b))
	a, b = 2, -5
	fmt.Println(getSum(a, b))
}

/*
output:
5
-3
*/
