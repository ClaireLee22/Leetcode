class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        maxArea = 0
        for i in range(len(height)):
            for j in range(i+1, len(height)):
                area = (j - i) * min(height[i], height[j])
                maxArea = max(maxArea, area)

        return maxArea


if __name__ == "__main__":
    height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
    solution = Solution()
    print(solution.maxArea(height))

# output: 49