package main

import (
	"fmt"
	"math"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
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

	prev := math.MaxInt32
	for i := pos; i < len(candidates); i++ {
		currentNum := candidates[i]
		if currentNum == prev {
			fmt.Printf("skip %d\n", currentNum)
			continue
		}

		currentCombination = append(currentCombination, currentNum)
		backtrack(candidates, target-currentNum, i+1, currentCombination, result)
		currentCombination = currentCombination[:len(currentCombination)-1]
		prev = currentNum
	}
}

func main() {
	candidates, target := []int{2, 3, 6, 7, 2}, 7
	fmt.Println("combinations:", combinationSum2(candidates, target))
}

/*
outout:
new combination: [2 2 3]
skip 2
new combination: [7]
combinations: [[2 2 3] [7]]
*/
