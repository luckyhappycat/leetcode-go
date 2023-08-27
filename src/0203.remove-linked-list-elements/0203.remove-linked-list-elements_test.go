package leetcode

import (
	"testing"

	"github.com/luckyhappycat/leetcode-go/structures"
)

type Param struct {
	head *ListNode
	val  int
}

type Result struct {
	ret *ListNode
}

type Case struct {
	param  Param
	result Result
}

func TestRemoveElements(t *testing.T) {
	cases := []Case{
		{
			param: Param{
				head: structures.InitListNode([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			result: Result{
				ret: structures.InitListNode([]int{1, 2, 3, 4, 5}),
			},
		},
		{
			param: Param{
				head: structures.InitListNode([]int{}),
				val:  1,
			},
			result: Result{
				ret: structures.InitListNode([]int{}),
			},
		},
		{
			param: Param{
				head: structures.InitListNode([]int{7, 7, 7, 7}),
				val:  7,
			},
			result: Result{
				ret: structures.InitListNode([]int{}),
			},
		},
	}
	for _, c := range cases {
		structures.PrintListNode(c.param.head)
		structures.PrintListNode(c.result.ret)
		got := removeElements(c.param.head, c.param.val)
		structures.PrintListNode(got)
		if !structures.EqualListNode(got, c.result.ret) {
			t.Errorf("removeElements(%v, %v) == %v, want %v", c.param.head, c.param.val, got, c.result.ret)
		}
	}
}
