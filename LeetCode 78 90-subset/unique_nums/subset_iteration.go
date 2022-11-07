package main

import "fmt"

func subsets(nums []int) [][]int {
	subsets := [][]int{{}}

	for _, num := range nums {
		length := len(subsets)
		for i := 0; i < length; i++ {
			newSubset := []int{}
			newSubset = append([]int{num}, subsets[i]...)
			subsets = append(subsets, newSubset)
		}
	}

	return subsets
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("subsets", subsets((nums)))
}

// output: subsets [[] [1] [2] [2 1] [3] [3 1] [3 2] [3 2 1]]
