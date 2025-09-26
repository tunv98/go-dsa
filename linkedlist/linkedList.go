package linkedlist

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func NewNode(val int) *ListNode {
	return &ListNode{
		val:  val,
		next: nil,
	}
}

func printOut(head *ListNode) {
	for head != nil {
		fmt.Print(head.val, " ")
		head = head.next
	}
	println()
}

func insertAtHead(head *ListNode, val int) *ListNode {
	newNode := NewNode(val)
	newNode.next = head
	return newNode
}

func insertAtTail(head *ListNode, val int) *ListNode {
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = NewNode(val)
	return head
}

func deleteHead(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	head = head.next
	return head
}

func deleteTail(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return nil
	}
	cur := head
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil
	return head
}

// list is sorted
func removeDuplicate(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.next != nil {
		if cur.val == cur.next.val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
	return head
}

func main() {
	list := NewNode(1)
	list.next = NewNode(2)
	list.next.next = NewNode(3)
	printOut(list)
	list = insertAtHead(list, 0)
	printOut(list)
	list = insertAtTail(list, 4)
	printOut(list)
	list = deleteHead(list)
	printOut(list)
	list = deleteTail(list)
	printOut(list)
	//-----------
	list = insertAtTail(list, 3) // 1 2 3 3
	list = insertAtTail(list, 4) // 1 2 3 3 4
	list = removeDuplicate(list)
	printOut(list)
}
