package leetcode

import (
	"github.com/luckyhappycat/leetcode-go/structures"
)

type ListNode = structures.ListNode

// https://leetcode-cn.com/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	for cur := dummyHead; cur != nil && cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
