package main

import "fmt"

// 给你一个单链表的头节点 head ，请你判断该链表是否为 回文链表 。如果是，返回 true ；否则，返回 false
type ListNode struct {
	Val  int
	Next *ListNode
}

//func isPalindrome(head *ListNode) bool {
//	l := make([]int, 0)
//	for ; head != nil; head = head.Next {
//		l = append(l, head.Val)
//	}
//	lens := len(l)
//	for i := 0; i < lens/2; i++ {
//		if l[i] != l[lens-1-i] {
//			return false
//		}
//	}
//	return true
//}

func isPalindrome(head *ListNode) bool {
	frontHead := head
	var recursion func(node *ListNode) bool
	recursion = func(h *ListNode) bool {
		if h != nil {
			if !recursion(h.Next) {
				return false
			}
			if frontHead.Val != h.Val {
				return false
			}
			frontHead = frontHead.Next
		}
		return true
	}
	return recursion(head)
}

func main() {
	a := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
	}
	fmt.Println(isPalindrome(a))
}
