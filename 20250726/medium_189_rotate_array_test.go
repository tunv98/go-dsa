package _0250726

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// inplace, O(1) extra space

// Space: O(n)
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n                                 // if rotate many rounds
	temp := append(nums[n-k:], nums[:n-k]...) // allocates a new slice → O(n) time and space.
	copy(nums, temp)                          // O(n)
}

// Space: O(1)
// to void extra space -> rotate many times
func rotateOptSpace(nums []int, k int) {
	nums = []int{1, 2}
	n := len(nums)
	k = k % n
	if n == 1 || n == 0 {
		return
	}
	reserve := func(nums []int, l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	reserve(nums, 0, n-1)
	reserve(nums, 0, k-1)
	reserve(nums, k, n-1)
}

func Test_rotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, arr)
}

func Test_rotateOptSpace(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotateOptSpace(arr, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, arr)
}
