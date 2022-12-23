class Solution(object):
     def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        result = 0
        operand = 0
        operator = 1

        for char in s:
            if '0' <= char <= '9':
                operand = operand * 10 + int(char)
            elif char in "-+":
                result += operand * operator
                operand = 0
                # save the operator we currently met for the next evaluation
                operator = 1 if char == "+" else -1

        result += operator * operand
        
        return result


if __name__ == "__main__":
    solution = Solution()
    s = "10+4-6"
    print("result", solution.calculate(s))

"""
output:
('result', 8)
"""
