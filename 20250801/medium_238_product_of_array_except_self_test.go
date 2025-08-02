package _0250801

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Time: O(n), Space: O(1)
func productExceptSelf(nums []int) []int {
	multiple := 1
	for _, num := range nums {
		if num == 0 {
			continue
		}
		multiple *= num
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = multiple
			continue
		}
		nums[i] = multiple / nums[i]
	}
	return nums
}

// Time: O(n), Space: O(n)
// without using the division operation
func productExceptSelfOpt(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	leftProduct := 1
	for i := 0; i < n; i++ {
		answer[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return answer
}
func Test_productExceptSelf(t *testing.T) {
	assert.Equal(t, []int{24, 12, 8, 6}, productExceptSelf([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{24, 12, 8, 6}, productExceptSelfOpt([]int{1, 2, 3, 4}))
}
