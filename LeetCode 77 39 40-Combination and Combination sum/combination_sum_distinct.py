class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        result = []
        
        def backtrack(pos, target, currentCombination):
            if target == 0:
                print("new combination", currentCombination)
                result.append(currentCombination[::])
                return
            
            if pos == len(candidates) or target < 0:
                return
            
            for i in range(pos, len(candidates)):
                currentElement = candidates[i]
                currentCombination.append(currentElement)
                backtrack(i, target-currentElement, currentCombination)
                currentCombination.pop()
            
        backtrack(0, target, [])
        
        return result

if __name__ == "__main__":
    solution = Solution()
    candidates, target = [2, 3, 6, 7], 7
    print("combinations", solution.combinationSum(candidates, target))


"""
output:
('new combination', [2, 2, 3])
('new combination', [7])
('combinations', [[2, 2, 3], [7]])
"""