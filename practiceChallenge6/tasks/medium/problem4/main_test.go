package main

import "testing"

func TestSubsetSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{3, 34, 4, 12, 5, 2}, 9, true},
		{[]int{3, 34, 4, 12, 5, 2}, 30, false},
		{[]int{1, 2, 3}, 0, true},
		{[]int{2, 4}, 3, false},
		{[]int{5}, 5, true},
		{[]int{5}, 4, false},
		{[]int{}, 0, true},
		{[]int{}, 5, false},
		{[]int{1, 2, 3}, 6, true},  // whole set
		{[]int{1, 2, 3}, 7, false}, // exceeds total
	}
	for _, c := range cases {
		if got := SubsetSum(c.nums, c.target); got != c.want {
			t.Errorf("SubsetSum(%v, %d) = %v, want %v", c.nums, c.target, got, c.want)
		}
	}
}
