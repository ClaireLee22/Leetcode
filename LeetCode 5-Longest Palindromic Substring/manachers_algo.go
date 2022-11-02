package main

import "fmt"

func longestPalindrome(s string) string {
	newString := insertBogusCharacters(s, '|')
	palindromeRadii := make([]int, len(newString))
	longestPalindromeCenter := 0

	center := 0
	radius := 0
	for center < len(newString) {

		// expend outwards
		expanedRadius := radius + 1
		for center-expanedRadius >= 0 && center+expanedRadius < len(newString) && newString[center-expanedRadius] == newString[center+expanedRadius] {
			radius = expanedRadius
			expanedRadius = radius + 1
		}

		palindromeRadii[center] = radius
		if palindromeRadii[center] > palindromeRadii[longestPalindromeCenter] {
			longestPalindromeCenter = center
		}

		oldCenter := center
		oldRadius := radius
		oldCenterRightBorder := oldCenter + oldRadius
		center += 1
		radius = 0
		for center <= oldCenterRightBorder {
			mirroredCenter := 2*oldCenter - center
			maxMirroredRadius := oldCenterRightBorder - center
			if palindromeRadii[mirroredCenter] < maxMirroredRadius {
				palindromeRadii[center] = palindromeRadii[mirroredCenter]
				center += 1
			} else if palindromeRadii[mirroredCenter] > maxMirroredRadius {
				palindromeRadii[center] = maxMirroredRadius
				center += 1
			} else {
				radius = maxMirroredRadius
				break
			}
		}
	}

	lognestPalindrome := newString[longestPalindromeCenter-palindromeRadii[longestPalindromeCenter] : longestPalindromeCenter+palindromeRadii[longestPalindromeCenter]+1]
	fmt.Println("palindromeRadii:", palindromeRadii)
	return trimBogusCharacters(lognestPalindrome, '|')
}

func insertBogusCharacters(s string, bogusCharacter rune) string {
	newString := make([]rune, 2*len(s)+1)
	newString[0] = bogusCharacter
	i := 1
	for _, char := range s {
		newString[i] = char
		newString[i+1] = bogusCharacter
		i += 2
	}
	return string(newString)
}

func trimBogusCharacters(s string, bogusCharacter rune) string {
	trimmedString := []rune{}
	for _, char := range s {
		if char != bogusCharacter {
			trimmedString = append(trimmedString, char)
		}
	}
	return string(trimmedString)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "aabab"
	fmt.Println("the longest palindrome:", longestPalindrome(s))
}

/*
output:
palindromeRadii: [0 1 2 1 0 3 0 3 0 1 0]
the longest palindrome: aba
*/
