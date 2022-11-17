package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	backtrack(candidates, target, 0, []int{}, &result)

	return result
}

func backtrack(candidates []int, target int, pos int, currentCombination []int, result *[][]int) {
	if target == 0 {
		newCombination := append([]int{}, currentCombination...)
		fmt.Println("new combination:", newCombination)
		*result = append(*result, newCombination)
		return
	}

	if pos == len(candidates) || target < 0 {
		return
	}

	for i := pos; i < len(candidates); i++ {
		currentNum := candidates[i]
		currentCombination = append(currentCombination, currentNum)
		backtrack(candidates, target-currentNum, i, currentCombination, result)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}

func main() {
	candidates, target := []int{2, 3, 6, 7}, 7
	fmt.Println("combinations:", combinationSum(candidates, target))
}

/*
outout:

*/
