package _0250713

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// sorted in non-decreasing order -> result is the same
func sortedSquares(nums []int) []int {
	i, j := 0, len(nums)-1
	resp := make([]int, len(nums))
	pos := len(nums) - 1
	for i <= j {
		iSquare := square(nums[i])
		jSquare := square(nums[j])
		if iSquare >= jSquare {
			resp[pos] = iSquare
			i++
		} else {
			resp[pos] = jSquare
			j--
		}
		pos--
	}
	return resp
}

func square(num int) int {
	return num * num
}

func Test_sortedSquares(t *testing.T) {
	assert.Equal(t, sortedSquares([]int{-4, -1, 0, 3, 10}), []int{0, 1, 9, 16, 100})
	assert.Equal(t, sortedSquares([]int{-7, -3, 2, 3, 11}), []int{4, 9, 9, 49, 121})
	assert.Equal(t, sortedSquares([]int{-5, -3, -2, -1}), []int{1, 4, 9, 25})
}
