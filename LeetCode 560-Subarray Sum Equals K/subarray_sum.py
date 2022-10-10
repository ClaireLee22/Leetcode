# Time: O(n^2) | Spae: O(1)
class BruteForceSolution(object):
    def subarraySum(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        count = 0
        for i in range(len(nums)):
            currentSum = 0
            for j in range(i, len(nums)):
                currentSum += nums[j]
                if currentSum == k:
                    count += 1
        return count


# Time: O(n) | Spae: O(n)
class OptimalSolution(object):
    def subarraySum(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        prefixSums = {0: 1}
        count = 0
        currentSum = 0
        for i in range(len(nums)):
            currentSum += nums[i]
            diff = currentSum - k
            if diff in prefixSums:
                count += prefixSums[diff]
            
            if currentSum in prefixSums:
                prefixSums[currentSum] += 1
            else:
                prefixSums[currentSum] = 1
            print("prefixSums", prefixSums)
                    
        return count
                

if __name__ == "__main__":
    bs = BruteForceSolution()
    os = OptimalSolution()
    array = [-1, 2, 5, -3, -1, 1, 1]
    k = 2
    print("count(brute force approach", bs.subarraySum(array, k))
    print("count", os.subarraySum(array, k))

"""
output:
('count(brute force approach', 5)
('prefixSums', {0: 1, -1: 1})
('prefixSums', {0: 1, 1: 1, -1: 1})
('prefixSums', {0: 1, 1: 1, 6: 1, -1: 1})
('prefixSums', {0: 1, 1: 1, 3: 1, 6: 1, -1: 1})
('prefixSums', {0: 1, 1: 1, 2: 1, 3: 1, 6: 1, -1: 1})
('prefixSums', {0: 1, 1: 1, 2: 1, 3: 2, 6: 1, -1: 1})
('prefixSums', {0: 1, 1: 1, 2: 1, 3: 2, 4: 1, 6: 1, -1: 1})
('count', 5)
"""