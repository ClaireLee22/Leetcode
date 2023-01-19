package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxPathSum := root.Val
	dfs(root, &maxPathSum)
	return maxPathSum
}

func dfs(root *TreeNode, maxPathSum *int) int {
	if root == nil {
		return 0
	}

	leftMax := max(dfs(root.Left, maxPathSum), 0)
	rightMax := max(dfs(root.Right, maxPathSum), 0)

	*maxPathSum = max(*maxPathSum, root.Val+leftMax+rightMax)
	return root.Val + max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: -10}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 20}
	node4 := &TreeNode{Val: 15}
	node5 := &TreeNode{Val: 7}
	root.Left = node2
	root.Right = node3
	node3.Left = node4
	node3.Right = node5

	fmt.Println(maxPathSum(root))
}

// output: 42
