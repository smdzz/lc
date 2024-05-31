// 环形链表

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
}
