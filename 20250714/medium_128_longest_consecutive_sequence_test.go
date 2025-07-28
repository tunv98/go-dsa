package _0250714

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// O(n) | O(n)
func longestConsecutive(nums []int) int {
	hashSet := make(map[int]bool)
	for _, num := range nums {
		hashSet[num] = true // use set to check existed
	}
	maxLen := 0
	for num := range hashSet {
		if hashSet[num-1] { // just start when "start a consecutive"
			continue
		}
		counter := 1
		for next := num + 1; hashSet[next]; next++ {
			counter++
		}
		if counter > maxLen {
			maxLen = counter
		}
	}
	return maxLen
}

func Test_longestConsecutive(t *testing.T) {
	assert.Equal(t, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), 9)
	assert.Equal(t, longestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)
}
