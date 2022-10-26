package main

import "fmt"

func minDistance(str1 string, str2 string) int {
	distance := make([][]int, len(str2)+1)
	for i := 0; i < len(distance); i++ {
		distance[i] = make([]int, len(str1)+1)
		for c := 0; c < len(str1)+1; c++ {
			distance[0][c] = c
		}
	}

	for r := 1; r < len(str2)+1; r++ {
		distance[r][0] = distance[r-1][0] + 1
	}

	for r := 1; r < len(str2)+1; r++ {
		for c := 1; c < len(str1)+1; c++ {
			if str1[c-1] == str2[r-1] {
				distance[r][c] = distance[r-1][c-1]
			} else {
				distance[r][c] = 1 + min(distance[r][c-1],
					distance[r-1][c-1],
					distance[r-1][c])
			}
		}
	}
	fmt.Println("distance matrix:", distance)
	return distance[len(str2)][len(str1)]
}

func min(args ...int) int {
	current := args[0]
	for i := 1; i < len(args); i++ {
		if current > args[i] {
			current = args[i]
		}
	}
	return current
}

func main() {
	str1 := "horse"
	str2 := "ios"
	fmt.Println("Levenshtein distance:", minDistance(str1, str2))
}

/*
output:
distance matrix: [[0 1 2 3 4 5] [1 1 2 2 3 4] [2 2 1 2 3 4] [3 3 2 2 2 3]]
Levenshtein distance: 3
*/
