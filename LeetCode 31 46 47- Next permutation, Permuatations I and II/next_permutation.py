class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        i = n - 1
        # find the longest decreasing sequence
        while i > 0 and nums[i-1] >= nums[i]:
            i -= 1
        
        # edge case: the largest permutation
        if i == 0:
            nums[:] = nums[::-1]
            return

        print("the longest decreasing subsequence", nums[i:])

        # identify the first non-decreasing element
        firstNonDecreasingElement = i - 1
        print("first non-decreasing element", nums[firstNonDecreasingElement], "index", firstNonDecreasingElement)
        
        # traverse backward to find the first element greater than the first non-decreasing element
        for j in range(n-1, firstNonDecreasingElement, -1):
            if nums[j] > nums[firstNonDecreasingElement]:
                print("first element > first non-decreasing element", nums[j])
                # swap
                print("before swap", nums)
                self.swap(j, firstNonDecreasingElement, nums)
                print("after swap", nums)
                # reverse the longest decreasing subsequence
                nums[i:len(nums)] = nums[i:len(nums)][::-1]
                print("reverse the longest decreasing subsequence", nums)
                break
    
    def swap(self, i, j, array):
        array[i], array[j] = array[j], array[i]
        

if __name__ == "__main__":
    nums = [1, 5, 6, 4, 3, 2]
    solution = Solution()
    solution.nextPermutation(nums)

"""
output:
('the longest decreasing subsequence', [6, 4, 3, 2])
('first non-decreasing element', 5, 'index', 1)
('first element > first non-decreasing element', 6)
('before swap', [1, 5, 6, 4, 3, 2])
('after swap', [1, 6, 5, 4, 3, 2])
('reverse the longest decreasing subsequence', [1, 6, 2, 3, 4, 5])
"""