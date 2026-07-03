package main

import "testing"

func TestReverseString(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
		{"hello", "olleh"},
		{"Go!", "!oG"},
	}
	for _, c := range cases {
		if got := ReverseString(c.in); got != c.want {
			t.Errorf("ReverseString(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFirstIndex(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 5, -1},
		{[]int{}, 5, -1},
		{[]int{5}, 5, 0},
		{[]int{1, 2, 2}, 2, 1},
		{[]int{7, 8, 9}, 9, 2},
	}
	for _, c := range cases {
		if got := FirstIndex(c.nums, c.target); got != c.want {
			t.Errorf("FirstIndex(%v, %d) = %d, want %d", c.nums, c.target, got, c.want)
		}
	}
}

func TestAllEven(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{}, true},
		{[]int{2, 4, 6}, true},
		{[]int{2, 3, 4}, false},
		{[]int{1}, false},
		{[]int{0}, true},
		{[]int{2, 4, 5}, false},
	}
	for _, c := range cases {
		if got := AllEven(c.nums); got != c.want {
			t.Errorf("AllEven(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
