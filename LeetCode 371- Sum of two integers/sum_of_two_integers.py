class Solution(object):
    def getSum(self, a, b):
        """
        :type a: int
        :type b: int
        :rtype: int
        """
        mask = 0xffffffff
        a = a & mask
        while b:
            tmp = (a ^ b) & mask
            carry = ((a & b) << 1) & mask
            a = tmp
            b = carry
        
        # a is negative
        if (a >> 31) & 1:
            return ~(a ^ mask)
        return a

if __name__ == "__main__":
    solution = Solution()
    a, b = 2, 3
    print(solution.getSum(a, b))
    a, b = 2, -5
    print(solution.getSum(a, b))

"""
output:
5
-3
"""