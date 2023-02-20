class Solution:
    def rotate(self, matrix):
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        # transpose
        for r in range(n):
            for c in range(r, n):
                matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c] 
        
        # flip horizontally
        for r in range(n):
            for c in range(n//2):
                matrix[r][c], matrix[r][n-1-c] = matrix[r][n-1-c], matrix[r][c]


if __name__ == "__main__":
    matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9],
    ]

    solution = Solution()
    solution.rotate(matrix)
    print(matrix)

# output: [[7, 4, 1], [8, 5, 2], [9, 6, 3]]