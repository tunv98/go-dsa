package _0250715

import (
	"fmt"
	"testing"
)

// brute force - check all substrings
// Time: O(n3), Space O(1)
func longestPalindromeBruteforce(s string) string {
	isPalindrome := func(str string) bool {
		left, right := 0, len(str)-1
		for left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	maxPalindrome := string(s[0])
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			subStr := s[i:j] // index from i to j-1
			if isPalindrome(subStr) && len(subStr) > len(maxPalindrome) {
				maxPalindrome = subStr
			}
		}
	}
	return maxPalindrome
}

// expand around centers
// each element, consider it is center and expand 2 sides
// 2 cases: palindrome is even or odd
// time: O(n) | space: O(1)
func longestPalindromeExpandAroundCenter(s string) string {
	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}
	if len(s) == 0 {
		return ""
	}
	start, maxLen := 0, 1
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(i, i)   // with odd
		len2 := expandAroundCenter(i, i+1) // with even
		currentLen := len1
		if len2 > len1 {
			currentLen = len2
		}
		if currentLen > maxLen {
			maxLen = currentLen
			start = i - (currentLen-1)/2
		}
	}
	return s[start : start+maxLen]
}

func Test_longestPalindromeBruteforce(t *testing.T) {
	testCases := []string{
		"babad",   // Expected: "bab" or "aba"
		"cbbd",    // Expected: "bb"
		"a",       // Expected: "a"
		"ac",      // Expected: "a" or "c"
		"racecar", // Expected: "racecar"
		"abcdcba", // Expected: "abcdcba"
		"",        // Expected: ""
	}

	for _, test := range testCases {
		result := longestPalindromeBruteforce(test)
		fmt.Printf("Input: \"%s\" -> Output: \"%s\"\n", test, result)
	}
}

func Test_longestPalindromeExpandAroundCenter(t *testing.T) {
	testCases := []string{
		"babad",   // Expected: "bab" or "aba"
		"cbbd",    // Expected: "bb"
		"a",       // Expected: "a"
		"ac",      // Expected: "a" or "c"
		"racecar", // Expected: "racecar"
		"abcdcba", // Expected: "abcdcba"
		"",        // Expected: ""
	}

	for _, test := range testCases {
		result := longestPalindromeExpandAroundCenter(test)
		fmt.Printf("Input: \"%s\" -> Output: \"%s\"\n", test, result)
	}
}
