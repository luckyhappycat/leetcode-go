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

func TestDeleteDuplicats(t *testing.T) {
	cases := []Case{
		{
			param: Param{
				head: structures.InitListNode([]int{1, 1, 2}),
			},
			result: Result{
				want: structures.InitListNode([]int{1, 2}),
			},
		},
		// {
		// 	param: Param{
		// 		head: structures.InitListNode([]int{1, 1, 2, 3, 3}),
		// 	},
		// 	result: Result{
		// 		want: structures.InitListNode([]int{1, 2, 3}),
		// 	},
		// },
		// {
		// 	param: Param{
		// 		head: structures.InitListNode([]int{}),
		// 	},
		// 	result: Result{
		// 		want: structures.InitListNode([]int{}),
		// 	},
		// },
	}
	for _, c := range cases {
		structures.PrintListNode(c.param.head)
		structures.PrintListNode(c.result.want)
		got := deleteDuplicates(c.param.head)
		structures.PrintListNode(got)
		if !structures.EqualListNode(got, c.result.want) {
			t.Errorf("deleteDuplicates(%v) == %v, want %v", c.param.head, got, c.result.want)
		}
	}
}
