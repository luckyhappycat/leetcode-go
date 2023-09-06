package leetcode

import (
	"github.com/luckyhappycat/leetcode-go/structures"
)

type ListNode = structures.ListNode

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
