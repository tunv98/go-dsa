package _0250712

import (
	"fmt"
	"sort"
	"testing"
)

// nums[i] + nums[j] + nums[k] == 0
// ko được trùng kết quả

// O(n2)
func threeSum(nums []int) [][]int {
	var rs [][]int
	sort.Ints(nums)                    // O(n log n)
	for i := 0; i < len(nums)-2; i++ { // O(n)
		if i > 0 && nums[i] == nums[i-1] { //* //O(1)
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k { // mỗi cặp j-k duyệt tối đa O(n)
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				rs = append(rs, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] { //**
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return rs
}

func Test_threeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

/*
(*)
nums := []int{-4, -1, -1, 0, 1, 2}
Lần lặp đầu: i = 1, nums[i] = -1 -> Ta tìm được 1 bộ 3: [-1, -1, 2]
Lần lặp sau: i = 2, nums[i] = -1 (giống hệt nums[i-1])
=> Nếu ta không continue, ta sẽ lại tìm ra bộ 3 [-1, -1, 2] — trùng lặp
*/

/*
(**)
nums = [-2, -2, 0, 0, 2, 2]
i,j,k : 0,2,5 -> sum = -2 + 0 + 2 = 0
i,k,k: 0,3,4 -> sum = -2 + 0 + 2= 0 -> cần skip vì trùng
*/
