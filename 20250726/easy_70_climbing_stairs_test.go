package _0250726

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 2^n
func climbstairsRecursion(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbstairsRecursion(n-1) + climbstairsRecursion(n-2)
}

func climbstairsBottomUp(n int) int {
	if n == 1 {
		return 1
	}
	a, b := 1, 2 // a: ways to reach step 1, b: ways to reach step 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b // a: keep ways no of previous step (i-1), b: keep ways no of current step (i)
	}
	return b
}

func Test_climbStairs(t *testing.T) {
	assert.Equal(t, 5, climbstairsRecursion(4))
	assert.Equal(t, 5, climbstairsBottomUp(4))
}
