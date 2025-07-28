package _0250715

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestPalindrome(s string) int {
	charCounter := make(map[rune]int)
	for _, c := range s {
		charCounter[c]++
	}
	result := 0
	hasOdd := false
	for _, count := range charCounter {
		result += (count / 2) * 2
		if count%2 == 1 {
			hasOdd = true
		}
	}
	if hasOdd {
		result++
	}
	return result
}

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, longestPalindrome("abccccdd"), 7)
	assert.Equal(t, longestPalindrome("ccc"), 3)
}
