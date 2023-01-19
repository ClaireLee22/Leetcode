class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution(object):
    def maxPathSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        def dfs(root, maxPathSum):
            if root is None:
                return 0

            leftMax = max(dfs(root.left, maxPathSum), 0)
            rightMax = max(dfs(root.right, maxPathSum), 0)

            maxPathSum[0] = max(maxPathSum[0], root.val + leftMax + rightMax)
            return root.val + max(leftMax, rightMax)
        
        maxPathSum = [root.val]
        dfs(root, maxPathSum)
        return maxPathSum[0]


if __name__ == "__main__":
    root = TreeNode(-10)
    node2 = TreeNode(9)
    node3 = TreeNode(20)
    node4 = TreeNode(15)
    node5 = TreeNode(7)
    root.left = node2
    root.right = node3
    node3.left = node4
    node3.right = node5

    solution = Solution()
    print(solution.maxPathSum(root))

# output: 42


