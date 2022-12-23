class Solution(object):
     def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        result = 0
        operand = 0
        operator = 1 # "+"
        stack = []
        
        for char in s:
            if '0' <= char <= '9':
                operand = operand * 10 + int(char)
            elif char in "-+":
                result += operand * operator
                # reset operator
                operand = 0
                # update operator
                operator = 1 if char == '+' else -1
            elif char == "(":
                stack.append(result)
                stack.append(operator)
                # reset result and operator
                result = 0
                operator = 1
            elif char == ")":
                result += operand * operator
                result *= stack.pop()
                result += stack.pop()
                # reset operand
                operand = 0
            
        result += operand * operator
                
        return result

if __name__ == "__main__":
    solution = Solution()
    s = "10+(2-3+5)-6"
    print("result", solution.calculate(s))

"""
output:
('result', 8)
"""