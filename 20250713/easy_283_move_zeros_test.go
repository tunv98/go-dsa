package _0250713

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// move all 0 to end
// inplace
// maintain relative order of other elements
func moveZeroes(nums []int) {
	insertPos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != insertPos {
				nums[insertPos] = nums[i]
			}
			insertPos++
		}
	}
	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}

func Test_moveZeroes(t *testing.T) {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	assert.Equal(t, arr, []int{1, 3, 12, 0, 0})
}
