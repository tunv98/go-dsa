package _0250712

import (
	"fmt"
	"testing"
)

// 1-indexed: Mảng này tính chỉ số từ 1
// sort in non-decreasing order

func twoSum(numbers []int, target int) []int {
	rs := make(map[int]int)
	for i, number := range numbers {
		if value, existed := rs[number]; existed {
			return []int{value, i + 1}
		}
		rs[target-number] = i + 1
	}
	return nil
}

// Your solution must use only constant extra space
// Hint1: 2 pointer
// Hint2: array is sorted

func twoSumII(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

func Test_twoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})
	fmt.Println(twoSumII([]int{2, 7, 11, 15}, 9), []int{1, 2})
	fmt.Println(twoSumII([]int{2, 7, 11, 15}, 9), []int{1, 2})
	fmt.Println(twoSumII([]int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997}, 542), []int{24, 32})
}
