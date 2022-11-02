class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        currentLongestPalindrome = [0, 1] # [leftIdx, rightIdx(exlusive)]
        for i in range(1, len(s)):
            odd = self.getLongestPalindrome(s, i-1, i+1)
            even = self.getLongestPalindrome(s, i-1, i)
            print("longest odd-length palindrome", odd, "longest even-length palindrome", even)
            potentialLongestPalindrome = max(odd, even, key=lambda x: x[1]-x[0])
            currentLongestPalindrome = max(currentLongestPalindrome, 
                                           potentialLongestPalindrome, 
                                           key=lambda x: x[1]-x[0])
        
        return s[currentLongestPalindrome[0]: currentLongestPalindrome[1]]
    
    def getLongestPalindrome(self, s, leftIdx, rightIdx):
        while leftIdx >= 0 and rightIdx < len(s):
            if s[leftIdx] != s[rightIdx]:
                break
            leftIdx -= 1
            rightIdx += 1
        
         # at the end of iteration leftIdx is subtracted by extra 1 and rightIdx is plus by extra 1
         # rightIdx is exlusive: rightIdx -1 + 1 = rightIdx
        return [leftIdx+1, rightIdx]

if __name__ == "__main__":
    solution = Solution()
    s = "aabab"
    print("the longest palindrome:", solution.longestPalindrome(s))

"""
output:
('longest odd-length palindrome', [1, 2], 'longest even-length palindrome', [0, 2])
('longest odd-length palindrome', [1, 4], 'longest even-length palindrome', [2, 2])
('longest odd-length palindrome', [2, 5], 'longest even-length palindrome', [3, 3])
('longest odd-length palindrome', [4, 5], 'longest even-length palindrome', [4, 4])
('the longest palindrome:', 'aba')
"""