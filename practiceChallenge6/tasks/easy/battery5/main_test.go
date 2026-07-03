package main

import "testing"

func TestRepeat(t *testing.T) {
	cases := []struct {
		s    string
		n    int
		want string
	}{
		{"ab", 3, "ababab"},
		{"x", 0, ""},
		{"", 5, ""},
		{"go", 1, "go"},
		{"-", 4, "----"},
	}
	for _, c := range cases {
		if got := Repeat(c.s, c.n); got != c.want {
			t.Errorf("Repeat(%q, %d) = %q, want %q", c.s, c.n, got, c.want)
		}
	}
}

func TestCountChar(t *testing.T) {
	cases := []struct {
		s      string
		target byte
		want   int
	}{
		{"hello", 'l', 2},
		{"", 'a', 0},
		{"aaa", 'a', 3},
		{"abc", 'z', 0},
		{"mississippi", 's', 4},
	}
	for _, c := range cases {
		if got := CountChar(c.s, c.target); got != c.want {
			t.Errorf("CountChar(%q, %q) = %d, want %d", c.s, c.target, got, c.want)
		}
	}
}

func TestIsSorted(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2, 2, 3}, true},
		{[]int{3, 1, 2}, false},
		{[]int{5, 4}, false},
		{[]int{1, 2, 3, 4, 5}, true},
	}
	for _, c := range cases {
		if got := IsSorted(c.nums); got != c.want {
			t.Errorf("IsSorted(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
