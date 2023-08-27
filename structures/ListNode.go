package structures

import (
	"bytes"
	"log"
	"strconv"
)

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

func PrintListNode(head *ListNode) {
	var buffer bytes.Buffer
	for head != nil {
		buffer.WriteString(strconv.Itoa(head.Val))
		head = head.Next
		if head != nil {
			buffer.WriteString("->")
		}
	}
	log.Println(buffer.String())
}

func EqualListNode(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1 == nil {
			return false
		}
		if l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
