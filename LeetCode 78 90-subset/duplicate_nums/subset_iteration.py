class Solution(object):
    def subsetsWithDup(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        subsets = [[]]
        startIdx, endIdx = 0, 0
        
        for idx, num in enumerate(nums):
            startIdx = 0
            if idx > 0 and nums[idx] == nums[idx - 1]:
                startIdx = endIdx + 1
            endIdx = len(subsets) - 1
            
            for j in range(startIdx, endIdx+1):
                newSubset = subsets[j][::]
                newSubset.append(num)
                subsets.append(newSubset)
                
        return subsets

if __name__ == "__main__":
    solution = Solution()
    nums = [2, 1, 2]
    print("subsets with duplicate", solution.subsetsWithDup(nums))

# output: ('subsets with duplicate', [[], [1], [2], [1, 2], [2, 2], [1, 2, 2]])