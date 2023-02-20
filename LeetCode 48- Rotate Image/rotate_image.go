package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	// transpose
	for r := 0; r < n; r++ {
		for c := r; c < n; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	// flip horozontally
	for r := 0; r < n; r++ {
		for c := 0; c < n/2; c++ {
			matrix[r][c], matrix[r][n-1-c] = matrix[r][n-1-c], matrix[r][c]
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)
	fmt.Println(matrix)
}

// output: [[7 4 1] [8 5 2] [9 6 3]]
