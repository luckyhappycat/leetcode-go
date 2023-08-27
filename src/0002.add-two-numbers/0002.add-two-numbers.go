package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	v1, v2, carry, cur := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: (v1 + v2 + carry) % 10, Next: nil}
		cur = cur.Next

		carry = (v1 + v2 + carry) / 10
	}
	return head.Next
}
