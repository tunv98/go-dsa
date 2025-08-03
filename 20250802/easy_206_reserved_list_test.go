package _0250802

import "go-dsa/utils"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	var prev *utils.ListNode
	curr := head
	for curr != nil {
		next := curr.Next // lưu node tiếp theo
		curr.Next = prev  // đảo chiều con trỏ
		prev = curr       // tiến prev
		curr = next       // tiến curr
	}
	return prev // prev là head mới
}
