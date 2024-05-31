package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var point = &ListNode{Val: 0, Next: nil}
	head := point
	carry := false
	for l1 != nil || l2 != nil || carry {
		sum := 0
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val  = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		if carry {
			sum += 1
			carry = false
		}
		sum += l1Val + l2Val
		if sum > 9 {
			sum = sum - 10
			carry = true
		}
		point.Next = &ListNode{Val: sum, Next: nil}
		point = point.Next
	}
	return head.Next
}

func main() {
	l1 := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}
	fmt.Println(addTwoNumbers(l1, l2))
}
