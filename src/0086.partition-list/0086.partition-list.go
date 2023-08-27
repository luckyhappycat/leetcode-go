package leetcode

import (
	"github.com/luckyhappycat/leetcode-go/structures"
)

type ListNode = structures.ListNode

func partition(head *ListNode, x int) *ListNode {
	lessHead := &ListNode{Val: 0, Next: nil}
	moreHead := &ListNode{Val: 0, Next: nil}
	lessCur := lessHead
	moreCur := moreHead
	for head != nil {
		if head.Val < x {
			lessCur.Next = head
			lessCur = lessCur.Next
		} else {
			moreCur.Next = head
			moreCur = moreCur.Next
		}
		head = head.Next
	}
	moreCur.Next = nil
	lessCur.Next = moreHead.Next
	return lessHead.Next
}
