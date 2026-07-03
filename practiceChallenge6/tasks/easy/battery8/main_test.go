package main

import "testing"

func TestIsEven(t *testing.T) {
	cases := []struct {
		n    int
		want bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{4, true},
		{7, false},
		{10, true},
	}
	for _, c := range cases {
		if got := IsEven(c.n); got != c.want {
			t.Errorf("IsEven(%d) = %v, want %v", c.n, got, c.want)
		}
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	cases := []struct {
		n    int
		want bool
	}{
		{1, true},
		{2, true},
		{4, true},
		{8, true},
		{16, true},
		{3, false},
		{6, false},
		{12, false},
		{100, false},
	}
	for _, c := range cases {
		if got := IsPowerOfTwo(c.n); got != c.want {
			t.Errorf("IsPowerOfTwo(%d) = %v, want %v", c.n, got, c.want)
		}
	}
}

func TestCollatzSteps(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 1},
		{3, 7},
		{6, 8},
		{7, 16},
	}
	for _, c := range cases {
		if got := CollatzSteps(c.n); got != c.want {
			t.Errorf("CollatzSteps(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}
