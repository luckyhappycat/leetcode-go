package leetcode

import (
	"testing"

	"github.com/luckyhappycat/leetcode-go/structures"
)

type Param struct {
	l1 *ListNode
	l2 *ListNode
}

type Result struct {
	want *ListNode
}

type Case struct {
	param  Param
	result Result
}

func TestMergeTwoLists(t *testing.T) {
	cases := []Case{
		{
			param: Param{
				l1: structures.InitListNode([]int{1, 2, 4}),
				l2: structures.InitListNode([]int{1, 3, 4}),
			},
			result: Result{
				want: structures.InitListNode([]int{1, 1, 2, 3, 4, 4}),
			},
		},
		{
			param: Param{
				l1: structures.InitListNode([]int{}),
				l2: structures.InitListNode([]int{}),
			},
			result: Result{
				want: structures.InitListNode([]int{}),
			},
		},
		{
			param: Param{
				l1: structures.InitListNode([]int{}),
				l2: structures.InitListNode([]int{0}),
			},
			result: Result{
				want: structures.InitListNode([]int{0}),
			},
		},
	}
	for _, c := range cases {
		structures.PrintListNode(c.param.l1)
		structures.PrintListNode(c.param.l2)
		structures.PrintListNode(c.result.want)
		got := mergeTwoLists(c.param.l1, c.param.l2)
		structures.PrintListNode(got)
		if structures.EqualListNode(got, c.result.want) == false {
			t.Errorf("mergeTwoLists(%v, %v) == %v, want %v", c.param.l1, c.param.l2, got, c.result.want)
		}
	}
}
