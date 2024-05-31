# -*- coding: utf-8 -*-
# @File  : add_sum.py
# @Author: mengde
# @Date  : 2021/11/11

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
            res = point = ListNode()
            carry = False
            while True:
                new_node = ListNode()
                sum = 0
                if l1 or l2 or carry:
                    l1_num = l1.val if l1 else 0
                    l2_num = l2.val if l2 else 0
                    l1 = l1.next if l1 else None
                    l2 = l2.next if l2 else None
                    if carry:
                        sum +=  1
                        carry = False
                    sum +=l1_num + l2_num
                    if sum < 10:
                        new_node.val = sum
                    else:
                        new_node.val = sum - 10
                        carry = True
                    point.next  = new_node
                    point = point.next
                else:
                    break
            return res.next
