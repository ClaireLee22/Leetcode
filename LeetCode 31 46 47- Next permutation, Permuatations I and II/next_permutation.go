package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1

	// find the longest decreasing sequence
	for i > 0 && nums[i-1] >= nums[i] {
		i -= 1
	}

	// edge case: the largest permutation
	if i == 0 {
		reverse(nums)
		return
	}

	fmt.Println("the longest decreasing subsequence:", nums[i:])
	// identify the first non-decreasing element
	firstNonDecreasingElement := i - 1
	fmt.Println("first non-decreasing element:", nums[firstNonDecreasingElement], "index:", firstNonDecreasingElement)

	// traverse backward to find the first element greater than the first non-decreasing element
	for j := n - 1; j > firstNonDecreasingElement; j-- {
		if nums[j] > nums[firstNonDecreasingElement] {
			fmt.Println("first element > first non-decreasing element:", nums[j])
			// swap
			fmt.Println("before swap:", nums)
			swap(j, firstNonDecreasingElement, nums)
			fmt.Println("after swap:", nums)
			// reverse the longest decreasing subsequence
			reverse(nums[i:])
			fmt.Println("reverse the longest decreasing subsequence:", nums)
			break
		}
	}
}

func reverse(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	nums := []int{1, 5, 6, 4, 3, 2}
	nextPermutation(nums)
}

/*
output:
the longest decreasing subsequence: [6 4 3 2]
first non-decreasing element: 5 index: 1
first element > first non-decreasing element: 6
before swap: [1 5 6 4 3 2]
after swap: [1 6 5 4 3 2]
reverse the longest decreasing subsequence: [1 6 2 3 4 5]
*/
