package main

import "fmt"

type substring struct {
	leftIdx  int
	rightIdx int // exclisuive
}

func (ss substring) length() int {
	return ss.rightIdx - ss.leftIdx
}

func longestPalindrome(s string) string {
	currentLongestPalindrome := substring{0, 1}

	for i := 1; i < len(s); i++ {
		odd := getLongestPalindrome(s, i-1, i+1)
		even := getLongestPalindrome(s, i-1, i)
		fmt.Println("longest odd-length palindrome", odd, "longest even-length palindrome", even)
		potentialLongestPalindrome := odd
		if even.length() > odd.length() {
			potentialLongestPalindrome = even
		}

		if potentialLongestPalindrome.length() > currentLongestPalindrome.length() {
			currentLongestPalindrome = potentialLongestPalindrome
		}
	}

	return s[currentLongestPalindrome.leftIdx:currentLongestPalindrome.rightIdx]
}

func getLongestPalindrome(s string, leftIdx int, rightIdx int) substring {
	for leftIdx >= 0 && rightIdx < len(s) {
		if s[leftIdx] != s[rightIdx] {
			break
		}

		leftIdx -= 1
		rightIdx += 1
	}
	// at the end of iteration leftIdx is subtracted by extra 1 and rightIdx is plus by extra 1
	// rightIdx is exlusive: rightIdx -1 + 1 = rightIdx
	return substring{leftIdx: leftIdx + 1, rightIdx: rightIdx}
}

func main() {
	s := "aabab"
	fmt.Println("the longest palindrome:", longestPalindrome(s))
}

/*
output:
longest odd-length palindrome {1 2} longest even-length palindrome {0 2}
longest odd-length palindrome {1 4} longest even-length palindrome {2 2}
longest odd-length palindrome {2 5} longest even-length palindrome {3 3}
longest odd-length palindrome {4 5} longest even-length palindrome {4 4}
the longest palindrome: aba
*/
