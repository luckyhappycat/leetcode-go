package leetcode

import (
	"reflect"
	"testing"
)

type Param struct {
	nums   []int
	target int
}

type Result struct {
	want []int
}

type Case struct {
	params Param
	result Result
}

func TestTwoSum(t *testing.T) {
	cases := []Case{
		{
			params: Param{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			result: Result{
				want: []int{0, 1},
			},
		},
		{
			params: Param{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			result: Result{
				want: []int{1, 2},
			},
		},
		{
			params: Param{
				nums:   []int{3, 3},
				target: 6,
			},
			result: Result{
				want: []int{0, 1},
			},
		},
	}

	for _, c := range cases {
		got := twoSum(c.params.nums, c.params.target)
		if !reflect.DeepEqual(got, c.result.want) {
			t.Errorf("twoSum(%v, %d) == %v, want %v", c.params.nums, c.params.target, got, c.result.want)
		}
	}
}
