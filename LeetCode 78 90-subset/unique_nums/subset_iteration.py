class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        subsets = [[]]
        
        for num in nums:
            length = len(subsets)
            for i in range(length):
                subsets.append([num] + subsets[i])
        
        return subsets

if __name__ == "__main__":
    solution = Solution()
    nums = [1, 2, 3]
    print("subsets", solution.subsets(nums))

# output: ('subsets', [[], [1], [2], [2, 1], [3], [3, 1], [3, 2], [3, 2, 1]])