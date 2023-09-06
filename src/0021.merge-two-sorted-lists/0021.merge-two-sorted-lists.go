package leetcode

import (
	"github.com/luckyhappycat/leetcode-go/structures"
)

type ListNode = structures.ListNode

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: nil}
	cur := dummyHead
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}


// 递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}