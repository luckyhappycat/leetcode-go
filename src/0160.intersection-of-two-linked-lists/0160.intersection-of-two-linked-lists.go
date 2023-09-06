package leetcode

import (
	"github.com/luckyhappycat/leetcode-go/structures"
)

type ListNode = structures.ListNode

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != nil || pB != nil {
		if pA == pB {
			return pA
		}

		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}

	}

	return nil
}
