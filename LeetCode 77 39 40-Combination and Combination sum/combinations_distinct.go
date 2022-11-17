package main

import "fmt"

func combine(n int, k int) [][]int {
	combinations := [][]int{}
	backtrack(n, k, 1, []int{}, &combinations)

	return combinations
}

func backtrack(n int, k int, start int, currentCombination []int, combinations *[][]int) {
	if len(currentCombination) == k {
		// deep copy
		newCombination := append([]int{}, currentCombination...)
		fmt.Println("new combination:", newCombination)
		*combinations = append(*combinations, newCombination)
		return
	}

	for i := start; i <= n; i++ {
		currentCombination = append(currentCombination, i)
		backtrack(n, k, i+1, currentCombination, combinations)
		// clean
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}

func main() {
	n, k := 4, 3
	fmt.Println("combinations:", combine(n, k))
}

/* outout:
new combination: [1 2 3]
new combination: [1 2 4]
new combination: [1 3 4]
new combination: [2 3 4]
combinations: [[1 2 3] [1 2 4] [1 3 4] [2 3 4]]
*/
