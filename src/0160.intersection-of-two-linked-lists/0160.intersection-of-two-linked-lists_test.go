package leetcode

import (
	"testing"

	"github.com/luckyhappycat/leetcode-go/structures"
)

type Param struct {
	headA *ListNode
	headB *ListNode
}

type Result struct {
	ret *ListNode
}

type Case struct {
	param  Param
	result Result
}

func TestGetIntersectionNode(t *testing.T) {
	cases := []Case{
		{
			param: Param{
				headA: structures.InitListNode([]int{4, 1, 8, 4, 5}),
				headB: structures.InitListNode([]int{5, 6, 1, 8, 4, 5}),
			},
			result: Result{
				ret: structures.InitListNode([]int{8, 4, 5}),
			},
		},
		{
			param: Param{
				headA: structures.InitListNode([]int{1, 9, 1, 2, 4}),
				headB: structures.InitListNode([]int{3, 2, 4}),
			},
			result: Result{
				ret: structures.InitListNode([]int{2, 4}),
			},
		},
		{
			param: Param{
				headA: structures.InitListNode([]int{2, 6, 4}),
				headB: structures.InitListNode([]int{1, 5}),
			},
			result: Result{
				ret: structures.InitListNode([]int{}),
			},
		},
	}
	for _, c := range cases {
		structures.PrintListNode(c.param.headA)
		structures.PrintListNode(c.param.headB)
		got := getIntersectionNode(c.param.headA, c.param.headB)
		structures.PrintListNode(got)
		if !structures.EqualListNode(got, c.result.ret) {
			t.Errorf("getIntersectionNode(%v, %v) == %v, want %v", c.param.headA, c.param.headB, got, c.result.ret)
		}
	}
}
