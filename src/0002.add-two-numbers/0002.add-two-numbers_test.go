package leetcode

import (
	"reflect"
	"testing"
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
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			result: Result{
				want: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  8,
							Next: nil,
						},
					},
				},
			},
		},
		{
			param: Param{
				l1: &ListNode{
					Val:  0,
					Next: nil,
				},
				l2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			result: Result{
				want: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
		},
		{
			param: Param{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
			result: Result{
				want: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
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
