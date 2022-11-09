class Solution(object):
    def nextPermutation(self, nums):
        n = len(nums)
        i = n - 1
        while i > 0 and nums[i-1] >= nums[i]:
            i -= 1
            
        if i == 0:
            nums[:] = nums[::-1]
            return
        
        firstNonDecreasingElement = i - 1
        for j in range(n-1, firstNonDecreasingElement, -1):
            if nums[j] > nums[firstNonDecreasingElement]:
                self.swap(j, firstNonDecreasingElement, nums)
                nums[i:] = nums[i:][::-1]
                break
        
    def permute(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        originalNums = nums[:]
        result = [originalNums]
        while True:
            self.nextPermutation(nums)
            if originalNums == nums:
                break
            else:
                result.append(nums[:])
                
        return result
    
    def swap(self, i, j, array):
        array[i], array[j] = array[j], array[i]


if __name__ == "__main__":
    numsArrs = [[1, 2, 3], [1, 2, 2]]
    solution = Solution()
    for numsArr in numsArrs:
        print("input array", numsArr)
        print("permutations", solution.permute(numsArr))

"""
output:
('input array', [1, 2, 3])
('permutations', [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]])
('input array', [1, 2, 2])
('permutations', [[1, 2, 2], [2, 1, 2], [2, 2, 1]])
"""