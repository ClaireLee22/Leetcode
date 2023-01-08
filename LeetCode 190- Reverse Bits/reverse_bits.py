class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
        res = 0
        for i in range(32):
            bit = (n >> i) & 1
            res += (bit << (31-i))
        return res


if __name__ == "__main__":
    solution = Solution()
    num = 43261596
    print(solution.reverseBits(num))

# output: 964176192