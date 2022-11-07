package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	backtrack(nums, 0, []int{}, &result)
	return result
}

func backtrack(nums []int, idx int, subset []int, result *[][]int) {
	if idx == len(nums) {
		newSubset := append([]int{}, subset...)
		*result = append(*result, newSubset)
		return
	}

	// subsets include nums[idx]
	subset = append(subset, nums[idx])
	backtrack(nums, idx+1, subset, result)
	subset = subset[0 : len(subset)-1]

	// subsets exclude nums[idx]
	for idx+1 < len(nums) && nums[idx] == nums[idx+1] {
		idx += 1
	}
	backtrack(nums, idx+1, subset, result)
}

func main() {
	nums := []int{2, 1, 2}
	fmt.Println("subsets with duplicate", subsetsWithDup(nums))
}

// output: subsets with duplicate [[1 2 2] [1 2] [1] [2 2] [2] []]
