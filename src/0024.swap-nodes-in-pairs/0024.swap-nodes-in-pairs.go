package leetcode

import (
	"github.com/luckyhappycat/leetcode-go/structures"
)

type ListNode = structures.ListNode

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for cur := dummy; cur != nil && cur.Next != nil && cur.Next.Next != nil; {
		cur, cur.Next, cur.Next.Next, cur.Next.Next.Next = cur.Next, cur.Next.Next, cur.Next.Next.Next, cur.Next
	}
	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		cur.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		cur = node1
	}
	return dummy.Next
}
