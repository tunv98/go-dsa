package _0250801

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// o(n)
func searchRange(nums []int, target int) []int {
	l, r := -1, -1
	for i, v := range nums {
		if v == target {
			if l == -1 {
				l = i
			}
			r = i
		}
	}
	return []int{l, r}
}

// O(logn)
func searchRangeBinarySearch(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	first := -1
	for l <= r {
		mid := (r + l) / 2
		if nums[mid] == target {
			first = mid
			r = mid - 1
			continue
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return first
}

func findLast(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	last := -1
	for l <= r {
		mid := (r + l) / 2
		if nums[mid] == target {
			last = mid
			l = mid + 1
			continue
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return last
}

// nums sorted in non-decrease
func Test_searchRange(t *testing.T) {
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))

	assert.Equal(t, []int{3, 4}, searchRangeBinarySearch([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, []int{-1, -1}, searchRangeBinarySearch([]int{5, 7, 7, 8, 8, 10}, 6))
}
