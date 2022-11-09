package main

import "fmt"

func permute(nums []int) [][]int {
	originalNums := append([]int{}, nums...)
	result := [][]int{originalNums}
	for {
		nextPermutation(nums)
		if isEqual(originalNums, nums) {
			break
		} else {
			newPerm := append([]int{}, nums...)
			result = append(result, newPerm)
		}
	}

	return result
}

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i -= 1
	}

	if i == 0 {
		reverse(nums)
		return
	}

	firstNonDecreasingElement := i - 1
	for j := n - 1; j > firstNonDecreasingElement; j-- {
		if nums[j] > nums[firstNonDecreasingElement] {
			swap(j, firstNonDecreasingElement, nums)
			reverse(nums[i:])
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

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	numsArrs := [][]int{{1, 2, 3}, {1, 2, 2}}
	for _, numsArr := range numsArrs {
		fmt.Println("input array:", numsArr)
		fmt.Println("permutations:", permute((numsArr)))
	}
}

/*
output:
input array: [1 2 3]
permutations: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
input array: [1 2 2]
permutations: [[1 2 2] [2 1 2] [2 2 1]]
*/
