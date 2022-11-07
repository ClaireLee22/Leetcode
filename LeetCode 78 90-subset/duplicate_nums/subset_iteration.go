package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	subsets := [][]int{{}}

	startIdx, endIdx := 0, 0
	for idx, num := range nums {
		startIdx = 0
		if idx > 0 && nums[idx] == nums[idx-1] {
			startIdx = endIdx + 1
		}
		endIdx = len(subsets) - 1

		for j := startIdx; j < endIdx+1; j++ {
			newSubset := append([]int{num}, subsets[j]...)
			subsets = append(subsets, newSubset)
		}
	}

	return subsets
}

func main() {
	nums := []int{2, 1, 2}
	fmt.Println("subsets with duplicate", subsetsWithDup(nums))
}

// output: subsets with duplicate [[] [1] [2] [2 1] [2 2] [2 2 1]]
