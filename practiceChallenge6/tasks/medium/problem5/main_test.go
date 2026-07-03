package main

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		sorted []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 7, 9}, 5, 2},
		{[]int{1, 3, 5, 7, 9}, 1, 0},
		{[]int{1, 3, 5, 7, 9}, 9, 4},
		{[]int{1, 3, 5, 7, 9}, 4, -1},
		{[]int{1, 3, 5, 7, 9}, 0, -1},
		{[]int{1, 3, 5, 7, 9}, 10, -1},
		{[]int{}, 5, -1},
		{[]int{42}, 42, 0},
		{[]int{42}, 7, -1},
		{[]int{2, 4, 6, 8, 10, 12}, 10, 4},
		{[]int{2, 4, 6, 8, 10, 12}, 2, 0},
		{[]int{2, 4, 6, 8, 10, 12}, 12, 5},
	}
	for _, c := range cases {
		if got := BinarySearch(c.sorted, c.target); got != c.want {
			t.Errorf("BinarySearch(%v, %d) = %d, want %d", c.sorted, c.target, got, c.want)
		}
	}
}
