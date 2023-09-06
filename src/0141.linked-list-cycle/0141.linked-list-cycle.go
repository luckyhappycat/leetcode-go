package leetcode

import (
	"github.com/luckyhappycat/leetcode-go/structures"
)

type ListNode = structures.ListNode

// https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
}
