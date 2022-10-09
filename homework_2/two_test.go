package main

import (
	"sort"
	"testing"
)

var testIntersection = []struct {
	a      []int
	b      []int
	result []int
}{
	{[]int{23, 3, 1, 2}, []int{6, 2, 4, 23}, []int{2, 23}},
	{[]int{1, 1, 1}, []int{1, 1, 1, 1}, []int{1, 1, 1}},
	{[]int{50, 2, 6, 8, 10, 4}, []int{4, 5, 6, 7, 2}, []int{2, 4, 6}},
}

func TestIntersection(t *testing.T) {
	for _, v := range testIntersection {
		result := Intersection(v.a, v.b)
		sort.Ints(result)

		if !Equal(result, v.result) {
			t.Error(
				"For", v.a,
				"and", v.b,
				"expected", v.result,
				"got", result,
			)
		}
	}
}
