package main

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	cases := []struct {
		name string
		in   []any
		want []int
	}{
		{"flat", []any{1, 2, 3}, []int{1, 2, 3}},
		{"one level", []any{1, []any{2, 3}, 4}, []int{1, 2, 3, 4}},
		{"deeply left", []any{[]any{[]any{1}}, 2}, []int{1, 2}},
		{"mixed depths",
			[]any{[]any{1, 2}, []any{3, []any{4, 5}}},
			[]int{1, 2, 3, 4, 5}},
		{"order preserved across nesting",
			[]any{1, []any{2, []any{3, 4}, 5}, 6},
			[]int{1, 2, 3, 4, 5, 6}},
	}
	for _, c := range cases {
		if got := Flatten(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: Flatten(%v) = %v, want %v", c.name, c.in, got, c.want)
		}
	}
}

func TestFlattenEmpty(t *testing.T) {
	empties := [][]any{
		{},
		{[]any{}},
		{[]any{}, []any{[]any{}}},
	}
	for _, in := range empties {
		if got := Flatten(in); len(got) != 0 {
			t.Errorf("Flatten(%v) = %v, want empty", in, got)
		}
	}
}
