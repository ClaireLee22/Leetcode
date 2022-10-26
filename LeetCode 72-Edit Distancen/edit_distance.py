class Solution(object):
    def minDistance(self, string1, string2):
        """
        :type string1: str
        :type string2: str
        :rtype: int
        """
        distance = [[c for c in range(len(string1)+1)] for r in range(len(string2)+1)]
        
        for r in range(1, len(string2)+1):
            distance[r][0] = distance[r-1][0] + 1
        
        for r in range(1, len(string2)+1):
            for c in range(1, len(string1)+1):
                if string1[c-1] == string2[r-1]:
                    distance[r][c] = distance[r-1][c-1]
                else:
                    distance[r][c] = 1 + min(distance[r][c-1], 
                                             distance[r-1][c-1], 
                                             distance[r-1][c])
        print("distance matrix", distance)
        return distance[-1][-1]

if __name__ == "__main__":
    solution = Solution()
    string1 = "horse"
    string2 = "ios"
    print("Levenshtein distance", solution.minDistance(string1, string2))

"""
output:
('distance matrix', [[0, 1, 2, 3, 4, 5], [1, 1, 2, 2, 3, 4], [2, 2, 1, 2, 3, 4], [3, 3, 2, 2, 2, 3]])
('Levenshtein distance', 3)
"""