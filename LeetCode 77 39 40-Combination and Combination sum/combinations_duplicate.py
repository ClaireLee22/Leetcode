class Solution:
    def combine(self, elements, k):
        elements.sort()
        combinations = []
        
        def backtrack(pos, currentCombination):
            if len(currentCombination) == k:
                # copy and append
                print("new combination", currentCombination)
                combinations.append(currentCombination[::])
                return
            
            prev = None
            for i in range(pos, len(elements)):
                currentElement = elements[i]
                if prev == currentElement:
                    print("skip", currentElement)
                    continue
                currentCombination.append(currentElement)
                backtrack(i+1, currentCombination)
                # clean
                currentCombination.pop()
                prev = currentElement
            
        backtrack(0, [])
        return combinations


if __name__ == "__main__":
    solution = Solution()
    elements = [1, 2, 3, 2]
    k = 3
    print("combinations", solution.combine(elements, k))


"""
output:
('new combination', [1, 2, 2])
('new combination', [1, 2, 3])
('skip', 2)
('new combination', [2, 2, 3])
('skip', 2)
('combinations', [[1, 2, 2], [1, 2, 3], [2, 2, 3]])
"""