package main

import "testing"

func TestContains(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{}, 5, false},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 5, false},
		{[]int{5}, 5, true},
		{[]int{1, 2, 3}, 1, true},
		{[]int{1, 2, 3}, 3, true},
	}
	for _, c := range cases {
		if got := Contains(c.nums, c.target); got != c.want {
			t.Errorf("Contains(%v, %d) = %v, want %v", c.nums, c.target, got, c.want)
		}
	}
}

func TestMaxSlice(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{5}, 5},
		{[]int{1, 2, 3}, 3},
		{[]int{3, 1, 2}, 3},
		{[]int{-5, -2, -9}, -2},
		{[]int{7, 7, 7}, 7},
	}
	for _, c := range cases {
		if got := MaxSlice(c.nums); got != c.want {
			t.Errorf("MaxSlice(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}

func TestCountOccurrences(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{}, 2, 0},
		{[]int{1, 2, 2, 3, 2}, 2, 3},
		{[]int{1, 2, 3}, 5, 0},
		{[]int{7, 7, 7}, 7, 3},
	}
	for _, c := range cases {
		if got := CountOccurrences(c.nums, c.target); got != c.want {
			t.Errorf("CountOccurrences(%v, %d) = %d, want %d", c.nums, c.target, got, c.want)
		}
	}
}
