package main

import (
	"reflect"
	"sort"
	"testing"
)

func sortedEqual(a, b []string) bool {
	a = append([]string(nil), a...)
	b = append([]string(nil), b...)
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}

func TestGenerateParentheses(t *testing.T) {
	cases := []struct {
		n    int
		want []string
	}{
		{0, []string{""}},
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, c := range cases {
		if got := GenerateParentheses(c.n); !sortedEqual(got, c.want) {
			t.Errorf("GenerateParentheses(%d) = %v, want (any order) %v", c.n, got, c.want)
		}
	}
}

func TestGenerateParenthesesCount(t *testing.T) {
	// count must equal the n-th Catalan number
	catalan := []int{1, 1, 2, 5, 14, 42} // n = 0..5
	for n := 0; n <= 5; n++ {
		if got := len(GenerateParentheses(n)); got != catalan[n] {
			t.Errorf("len(GenerateParentheses(%d)) = %d, want %d", n, got, catalan[n])
		}
	}
}
