package leetcode

import (
	"reflect"
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

func TestAddTwoNumbers(t *testing.T) {
	cases := []Case{
		{
			param: Param{
				l1: structures.InitListNode([]int{2, 4, 3}),
				l2: structures.InitListNode([]int{5, 6, 4}),
			},
			result: Result{
				want: structures.InitListNode([]int{7, 0, 8}),
			},
		},
		{
			param: Param{
				l1: structures.InitListNode([]int{0}),
				l2: structures.InitListNode([]int{0}),
			},
			result: Result{
				want: structures.InitListNode([]int{0}),
			},
		},
		{
			param: Param{
				l1: structures.InitListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: structures.InitListNode([]int{9, 9, 9, 9}),
			},
			result: Result{
				want: structures.InitListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
			},
		},
	}

	for _, c := range cases {
		got := addTwoNumbers(c.param.l1, c.param.l2)
		if !reflect.DeepEqual(got, c.result.want) {
			t.Errorf("addTwoNumbers(%v, %v) == %v, want %v", c.param.l1, c.param.l2, got, c.result.want)
		}
	}
}
