from collections import deque
class Solution(object):
     def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        array = deque(s)
        return self.calculateHelper(array)
    
     def calculateHelper(self, array):
        operand = 0
        operator = 1 # +
        stack = []
        
        while array:
            char = array.popleft()
            if "0" <= char <= "9":
                operand = operand * 10 + int(char)
            elif char in "-+":
                stack.append(operand*operator)
                operand = 0
                operator = 1 if char == "+" else -1
            elif char == "(":
                operand = self.calculateHelper(array)
            elif char == ")":
                break
            
            print(array)
        
        return sum(stack) + operand*operator

if __name__ == "__main__":
    solution = Solution()
    s = "10+(2-3+5)-6"
    print("result", solution.calculate(s))

"""
output:
('result', 8)
"""