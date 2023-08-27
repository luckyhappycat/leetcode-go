package leetcode

import (
	"reflect"
	"testing"

	"github.com/luckyhappycat/leetcode-go/structures"
)

type Param struct {
	head *ListNode
	n    int
}

type Result struct {
	ret *ListNode
}

type Case struct {
	param  Param
	result Result
}

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	cases := []Case{
		{
			param: Param{
				head: structures.InitListNode([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			result: Result{
				ret: structures.InitListNode([]int{1, 2, 3, 5}),
			},
		},
		{
			param: Param{
				head: structures.InitListNode([]int{1}),
				n:    1,
			},
			result: Result{
				ret: nil,
			},
		},
		{
			param: Param{
				head: structures.InitListNode([]int{1, 2}),
				n:    1,
			},
			result: Result{
				ret: structures.InitListNode([]int{1}),
			},
		},
	}
	for _, c := range cases {
		ret := removeNthFromEnd(c.param.head, c.param.n)
		if !reflect.DeepEqual(ret, c.result.ret) {
			t.Fatalf("case %v failed\nactual: %v, expect: %v\n", c.param, ret, c.result.ret)
		}
	}

}
