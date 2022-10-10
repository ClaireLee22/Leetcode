package main

import "fmt"

// Time: O(n^2) | Spae: O(1)
func bruteForceSubarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		currentSum := 0
		for j := i; j < len(nums); j++ {
			currentSum += nums[j]
			if currentSum == k {
				count += 1
			}
		}
	}

	return count
}

// Time: O(n) | Spae: O(n)
func subarraySum(nums []int, k int) int {
	prefixSums := map[int]int{0: 1}
	count := 0
	currentSum := 0

	for _, num := range nums {
		currentSum += num
		diff := currentSum - k
		if _, found := prefixSums[diff]; found {
			count += prefixSums[diff]
		}

		if _, found := prefixSums[currentSum]; found {
			prefixSums[currentSum] += 1
		} else {
			prefixSums[currentSum] = 1
		}
		fmt.Println("prefixSums: ", prefixSums)
	}

	return count
}

func main() {
	array := []int{-1, 2, 5, -3, -1, 1, 1}
	k := 2
	fmt.Println("count(brute force approach:", bruteForceSubarraySum(array, k))
	fmt.Println("count:", subarraySum(array, k))
}

/*
output:
count(brute force approach: 5
prefixSums:  map[-1:1 0:1]
prefixSums:  map[-1:1 0:1 1:1]
prefixSums:  map[-1:1 0:1 1:1 6:1]
prefixSums:  map[-1:1 0:1 1:1 3:1 6:1]
prefixSums:  map[-1:1 0:1 1:1 2:1 3:1 6:1]
prefixSums:  map[-1:1 0:1 1:1 2:1 3:2 6:1]
prefixSums:  map[-1:1 0:1 1:1 2:1 3:2 4:1 6:1]
count: 5
*/
