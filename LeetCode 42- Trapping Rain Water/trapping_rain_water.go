package main

import "fmt"

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	water := 0
	for left < right {
		if leftMax < rightMax {
			left += 1
			leftMax = max(leftMax, height[left])
			water += leftMax - height[left]
		} else {
			right -= 1
			rightMax = max(rightMax, height[right])
			water += rightMax - height[right]
		}
	}
	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(heights))
}

// output: 6
