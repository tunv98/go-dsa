package _0250801

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxLength := -1
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > maxLength {
			maxLength = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxLength
}

/*
[1, 8, 6, 2, 5, 4, 8, 3, 7]
 ↑                          ↑
left=0                   right=8
height[0]=1 < height[8]=7

Diện tích = min(1,7) * 8 = 1 * 8 = 8

Vì height[0] thấp hơn (luôn là min của 1 và height[k] -> 1),
mọi cặp (0, k) với k < 8
sẽ có diện tích ≤ 1 * (k-0) < 8
→ Loại bỏ được và di chuyển left++
*/

func Test_maxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 7, maxArea([]int{8, 7, 2, 1}))
}
