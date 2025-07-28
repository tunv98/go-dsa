package _0250724

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getConcatenation(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}

func Test_getConcatenation(t *testing.T) {
	assert.Equal(t, getConcatenation([]int{1, 2, 3, 4, 5}), []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5})
}
