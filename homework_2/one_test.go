package main

import "testing"

var testEqual = []struct {
	a      []int
	b      []int
	result bool
}{
	{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
	{[]int{1, 2, 3}, []int{3, 2, 1}, false},
	{[]int{10, 11, 12}, []int{10, 11, 12}, true},
}

func TestEqual(t *testing.T) {
	for _, v := range testEqual {
		result := Equal(v.a, v.b)
		if result != v.result {
			t.Error(
				"For", v.a,
				"and", v.b,
				"expected", v.result,
				"got", result,
			)
		}
	}
}

var testTwoSum = []struct {
	nums   []int
	target int
	result []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, v := range testTwoSum {
		result := TwoSum(v.nums, v.target)
		if !Equal(result, v.result) {
			t.Error(
				"For", v.nums,
				"and", v.target,
				"expected", v.result,
				"got", result,
			)
		}
	}
}
