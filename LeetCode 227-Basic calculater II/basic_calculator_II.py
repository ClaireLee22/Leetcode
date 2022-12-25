class Solution(object):
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        finalResult, tempResult, operand = 0, 0, 0
        operator = "+"
        
        for idx, char in enumerate(s):
            if '0' <= char <= '9':
                operand = operand * 10 + int(char)
            if idx == len(s) - 1 or char in "+-*/":
                if operator in "+-":
                    finalResult += tempResult
                    tempResult = operand if operator == "+" else -operand
                elif operator == "*":
                    tempResult *= operand
                elif operator == "/":
                    if tempResult < 0:
                        tempResult = -1 * (abs(tempResult)/operand)
                    else:
                        tempResult /= operand
            
                operand = 0
                operator = char
                
        
        finalResult += tempResult
        
        return finalResult

if __name__ == "__main__":
    solution = Solution()
    print("result", solution.calculate("10 + 2*3 - 6/3"))

"""
output:
('result', 14)
"""