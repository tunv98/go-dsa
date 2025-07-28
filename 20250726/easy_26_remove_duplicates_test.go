package _0250726

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// sorted in non-decreasing order
// remove the duplicates in-place -> each unique element only one
func removeDuplicates(nums []int) int {
	i := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[j] == nums[i] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}
func Test_removeDuplicates(t *testing.T) {
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums := []int{0, 1, 2, 3, 4}

	k := removeDuplicates(nums)
	assert.Equal(t, len(expectedNums), k)
	for i := 0; i < k; i++ {
		assert.Equal(t, expectedNums[i], nums[i])
	}
}
