package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, 0, []int{}, &result)
	return result
}

func backtrack(nums []int, idx int, subset []int, result *[][]int) {
	if idx == len(nums) {
		newSubset := append([]int{}, subset...) // copy
		*result = append(*result, newSubset)
		return
	}

	// subsets include nums[idx]
	subset = append(subset, nums[idx])
	backtrack(nums, idx+1, subset, result)
	subset = subset[:len(subset)-1]

	// subsets exclude nums[idx]
	backtrack(nums, idx+1, subset, result)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("subsets", subsets((nums)))
}

// output: subsets [[1 2 3] [1 2] [1 3] [1] [2 3] [2] [3] []]
