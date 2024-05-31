//  相交链表

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA == nil || headB == nil {
//		return nil
//	}
//	tmp := map[*ListNode]struct{}{}
//	pa := headA
//	for {
//		tmp[pa] = struct{}{}
//		pa = pa.Next
//		if pa == nil {
//			break
//		}
//	}
//	pb := headB
//	for {
//		if _, ok := tmp[pb]; ok {
//			return pb
//		}
//		pb = pb.Next
//		if pb == nil {
//			break
//		}
//
//	}
//	return nil
//}

// 假设两个链表相交,第一个链表长度为m, 第二个链表长度为n,第一个链表相交之前为a,第二个链表相交之前为b,相交之后的距离为c,那么a + c = m; b + c = n
// 那么a + c + b = b + c + a
// 所以当第一个链表走到尽头时去走第二个链表的路,第二个链表走到尽头时去走第一个链表的路 走完相同的路有值就是相交,都为nil就是不相交
// 走到尽头见不到你，于是走过你来时的路，等到相遇时才发现，你也走过我来时的路。
// 对的人错过了还是会相遇， 错的人相遇了也是NULL

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
