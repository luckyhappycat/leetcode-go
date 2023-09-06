package leetcode

import (
	"github.com/luckyhappycat/leetcode-go/structures"
)

type ListNode = structures.ListNode

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	fast, slow := head, dummyHead
	for fast != nil {
		if n <= 0 {
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
