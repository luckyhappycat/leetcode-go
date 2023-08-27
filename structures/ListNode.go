package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitListNode(arrs []int) *ListNode {
	if len(arrs) == 0 {
		return nil
	}
	head := &ListNode{Val: arrs[0]}
	cur := head
	for i := 1; i < len(arrs); i++ {
		cur.Next = &ListNode{Val: arrs[i]}
		cur = cur.Next
	}
	return head
}
