package _0250726

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func Test_plusOne(t *testing.T) {
	assert.Equal(t, []int{1, 0, 0, 0, 0}, plusOne([]int{9, 9, 9, 9}))
}
