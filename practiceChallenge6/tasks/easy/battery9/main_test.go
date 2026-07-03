package main

import (
	"reflect"
	"testing"
)

func TestRemoveChar(t *testing.T) {
	cases := []struct {
		s    string
		c    byte
		want string
	}{
		{"hello", 'l', "heo"},
		{"aaa", 'a', ""},
		{"abc", 'z', "abc"},
		{"", 'a', ""},
		{"banana", 'a', "bnn"},
	}
	for _, c := range cases {
		if got := RemoveChar(c.s, c.c); got != c.want {
			t.Errorf("RemoveChar(%q, %q) = %q, want %q", c.s, c.c, got, c.want)
		}
	}
}

func TestConcatSlice(t *testing.T) {
	cases := []struct {
		strs []string
		want string
	}{
		{[]string{}, ""},
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{"hello", " ", "world"}, "hello world"},
		{[]string{"solo"}, "solo"},
	}
	for _, c := range cases {
		if got := ConcatSlice(c.strs); got != c.want {
			t.Errorf("ConcatSlice(%v) = %q, want %q", c.strs, got, c.want)
		}
	}
}

func TestRange(t *testing.T) {
	cases := []struct {
		start, end int
		want       []int
	}{
		{1, 3, []int{1, 2, 3}},
		{5, 5, []int{5}},
		{0, 4, []int{0, 1, 2, 3, 4}},
	}
	for _, c := range cases {
		if got := Range(c.start, c.end); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Range(%d, %d) = %v, want %v", c.start, c.end, got, c.want)
		}
	}
	// start > end -> empty
	if got := Range(3, 2); len(got) != 0 {
		t.Errorf("Range(3, 2) = %v, want empty", got)
	}
}
