package leetcode

import (
	"testing"

	"github.com/luckyhappycat/leetcode-go/structures"
)

type Param struct {
	head *ListNode
}

type Result struct {
	want *ListNode
}

type Case struct {
	param  Param
	result Result
}

func TestSwapPairs(t *testing.T) {
	cases := []Case{
		{
			param: Param{
				head: structures.InitListNode([]int{1, 2, 3, 4}),
			},
			result: Result{
				want: structures.InitListNode([]int{2, 1, 4, 3}),
			},
		},
		{
			param: Param{
				head: structures.InitListNode([]int{}),
			},
			result: Result{
				want: structures.InitListNode([]int{}),
			},
		},
		{
			param: Param{
				head: structures.InitListNode([]int{1}),
			},
			result: Result{
				want: structures.InitListNode([]int{1}),
			},
		},
	}
	for _, c := range cases {
		structures.PrintListNode(c.param.head)
		structures.PrintListNode(c.result.want)
		got := swapPairs(c.param.head)
		structures.PrintListNode(got)
		if !structures.EqualListNode(got, c.result.want) {
			t.Errorf("swapPairs(%v) == %v, want %v", c.param.head, got, c.result.want)
		}
	}
}
