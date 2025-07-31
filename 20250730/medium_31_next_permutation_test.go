package _0250730

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// permutation = hoán vị
// inplace,constant extra memory
// • [1,2,3] → [1,3,2]
// • [3,2,1] → [1,2,3] (vì không còn cách sắp xếp nào lớn hơn, nên quay về nhỏ nhất)
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i := n - 2 // find i from right to left with nums[i] < nums[i+1]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1 // find j from right to left with nums[j] > nums[i]
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i] // swap
	}
	l, r := i+1, n-1 // swap segment i+1:end
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func Test_nextPermutation(t *testing.T) {
	arr1 := []int{1, 2, 5}
	nextPermutation(arr1)
	assert.Equal(t, arr1, []int{1, 5, 2})

	//-------
	arr2 := []int{3, 2, 1}
	nextPermutation(arr2)
	assert.Equal(t, arr2, []int{1, 2, 3})

	//-------
	arr3 := []int{1, 1}
	nextPermutation(arr3)
	assert.Equal(t, arr3, []int{1, 1})
}

/*
Các bước giải:
 1. Tìm điểm “giảm” đầu tiên từ phải sang trái
Duyệt từ cuối mảng về đầu, tìm chỉ số i sao cho nums[i] < nums[i+1].
Nếu không tìm thấy, tức là mảng đã là hoán vị lớn nhất, chỉ cần đảo ngược toàn bộ mảng.
 2. Tìm phần tử nhỏ nhất lớn hơn nums[i] ở bên phải
Duyệt từ cuối mảng về đầu, tìm chỉ số j sao cho nums[j] > nums[i].
 3. Hoán đổi nums[i] và nums[j]
 4. Đảo ngược đoạn từ i+1 đến hết mảng
Để đảm bảo phần phía sau là nhỏ nhất.
*/
