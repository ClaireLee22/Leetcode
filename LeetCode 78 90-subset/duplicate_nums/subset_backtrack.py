class Solution(object):
    def subsetsWithDup(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result = []
        nums.sort()
        
        def backtrack(idx, subset):
            if idx == len(nums):
                result.append(subset[::])
                return
            
            # subsets include nums[idx]
            subset.append(nums[idx])
            backtrack(idx+1, subset)
            subset.pop()
            
            # subsets exclude nums[idx]
            while idx + 1 < len(nums) and nums[idx] == nums[idx+1]:
                idx += 1
            backtrack(idx+1, subset)
        
        backtrack(0, [])
        return result
        
if __name__ == "__main__":
    solution = Solution()
    nums = [2, 1, 2]
    print("subsets with duplicate", solution.subsetsWithDup(nums))


# output: ('subsets with duplicate', [[1, 2, 2], [1, 2], [1], [2, 2], [2], []])