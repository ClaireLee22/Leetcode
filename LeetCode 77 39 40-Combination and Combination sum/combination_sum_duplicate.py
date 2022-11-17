class Solution(object):
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        candidates.sort()
        result = []
        
        def backtrack(pos, currentTarget, currentCombination):
            if currentTarget == 0:
                print("new combination", currentCombination)
                result.append(currentCombination[::])
                return
            
            if pos == len(candidates) or currentTarget < 0:
                return
            
            prev = None
            for i in range(pos, len(candidates)):
                currentNum = candidates[i]
                if currentNum == prev:
                    print("skip", currentNum)
                    continue
                
                currentCombination.append(currentNum)
                backtrack(i+1, currentTarget - currentNum, currentCombination)
                currentCombination.pop()
                prev = currentNum
            
        backtrack(0, target, [])
        
        return result
        
if __name__ == "__main__":
    solution = Solution()
    candidates, target = [2, 3, 6, 7, 2], 7
    print("combinations", solution.combinationSum2(candidates, target))

"""
output:
('new combination', [2, 2, 3])
('skip', 2)
('new combination', [7])
('combinations', [[2, 2, 3], [7]])
"""