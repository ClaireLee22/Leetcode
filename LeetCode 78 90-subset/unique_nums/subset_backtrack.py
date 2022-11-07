class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result = []
    
        def backtrack(idx, subset):
            if idx == len(nums):
                result.append(subset[::]) # copy
                return

            # sunsets includ nums[idx]
            subset.append(nums[idx])
            backtrack(idx+1, subset)
            subset.pop()

            # subsets exclude nums[idx]
            backtrack(idx+1, subset)

        backtrack(0, [])
        return result


if __name__ == "__main__":
    solution = Solution()
    nums = [1, 2, 3]
    print("subsets", solution.subsets(nums))

# output: ('subsets', [[1, 2, 3], [1, 2], [1, 3], [1], [2, 3], [2], [3], []])