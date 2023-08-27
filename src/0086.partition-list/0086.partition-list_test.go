package leetcode

import (
	"testing"

	"github.com/luckyhappycat/leetcode-go/structures"
)

type Param struct {
	head *ListNode
	x    int
}

type Result struct {
	want *ListNode
}

type Case struct {
	param  Param
	result Result
}

func TestPartition(t *testing.T) {
	caces := []Case{
		{
			param: Param{
				head: structures.InitListNode([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			result: Result{
				want: structures.InitListNode([]int{1, 2, 2, 4, 3, 5}),
			},
		},
		{
			param: Param{
				head: structures.InitListNode([]int{2, 1}),
				x:    2,
			},
			result: Result{
				want: structures.InitListNode([]int{1, 2}),
			},
		},
		{
			param: Param{
				head: structures.InitListNode([]int{}),
				x:    0,
			},
			result: Result{
				want: structures.InitListNode([]int{}),
			},
		},
	}
	for _, c := range caces {
		structures.PrintListNode(c.param.head)
		structures.PrintListNode(c.result.want)
		got := partition(c.param.head, c.param.x)
		structures.PrintListNode(got)
		if !structures.EqualListNode(got, c.result.want) {
			t.Errorf("partition(%v, %v) == %v, want %v", c.param.head, c.param.x, got, c.result.want)
		}
	}

}
