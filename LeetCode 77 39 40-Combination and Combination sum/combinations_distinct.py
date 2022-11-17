class Solution:
    def combine(self, n, k):
        combinations = []
        
        def backtrack(start, currentCombination):
            if len(currentCombination) == k:
                # copy and append
                print("new combination", currentCombination)
                combinations.append(currentCombination[::])
                return
            
            for i in range(start, n+1):
                currentCombination.append(i)
                backtrack(i+1, currentCombination)
                # clean
                currentCombination.pop()
            
        backtrack(1, [])
        return combinations


if __name__ == "__main__":
    solution = Solution()
    n, k = 4, 3
    print("combinations", solution.combine(n, k))


"""
output:
('new combination', [1, 2, 3])
('new combination', [1, 2, 4])
('new combination', [1, 3, 4])
('new combination', [2, 3, 4])
('combinations', [[1, 2, 3], [1, 2, 4], [1, 3, 4], [2, 3, 4]])
"""