//  反转链表?

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 在遍历链表时，将当前节点的 next 指针改为指向前一个节点。
//由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。
//在更改引用之前，还需要存储后一个节点。最后返回新的头引用。

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	newHead := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return newHead
//}

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
	read(reverseList(head))
}
