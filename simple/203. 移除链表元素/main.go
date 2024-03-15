package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
//func removeElements(head *ListNode, val int) *ListNode {
//	if head == nil {
//		return head
//	}
//	head.Next = removeElements(head.Next, val)
//	if head.Val == val {
//		return head.Next
//	}
//	return head
//}

// 遍历
func removeElements(head *ListNode, val int) *ListNode {
	temp := &ListNode{Next: head}
	tmp := temp
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return temp.Next
}

func read(head *ListNode) {
	if head == nil {
		return
	}
	for {
		fmt.Println(head.Val)
		head = head.Next
		if head == nil {
			return
		}
	}
}

func main() {
	//[1,2,6,3,4,5,6]
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}}
	a := removeElements(head, 6)
	read(a)
}
