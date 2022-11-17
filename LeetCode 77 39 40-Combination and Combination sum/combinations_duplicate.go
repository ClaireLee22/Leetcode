package main

import (
	"fmt"
	"math"
	"sort"
)

func combine(elements []int, k int) [][]int {
	sort.Ints(elements)
	combinations := [][]int{}
	backtrack(elements, k, 0, []int{}, &combinations)

	return combinations
}

func backtrack(elements []int, k int, pos int, currentCombination []int, combinations *[][]int) {
	if len(currentCombination) == k {
		// deep copy
		newCombination := append([]int{}, currentCombination...)
		fmt.Println("new combination:", newCombination)
		*combinations = append(*combinations, newCombination)
		return
	}

	prev := math.MaxInt32
	for i := pos; i < len(elements); i++ {
		currentElement := elements[i]
		if currentElement == prev {
			fmt.Printf("skip %d\n", currentElement)
			continue
		}
		currentCombination = append(currentCombination, currentElement)
		backtrack(elements, k, i+1, currentCombination, combinations)
		// clean
		currentCombination = currentCombination[:len(currentCombination)-1]
		prev = currentElement
	}
}

func main() {
	elements := []int{1, 2, 3, 2}
	k := 3
	fmt.Println("combinations:", combine(elements, k))
}

/*
outout:
new combination: [1 2 2]
new combination: [1 2 3]
skip 2
new combination: [2 2 3]
skip 2
combinations: [[1 2 2] [1 2 3] [2 2 3]]
*/
