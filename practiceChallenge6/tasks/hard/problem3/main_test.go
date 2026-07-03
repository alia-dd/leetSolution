package main

import "testing"

func TestWordBreak(t *testing.T) {
	cases := []struct {
		name string
		s    string
		dict []string
		want bool
	}{
		{"leetcode", "leetcode", []string{"leet", "code"}, true},
		{"reuse word", "applepenapple", []string{"apple", "pen"}, true},
		{"dead-end greedy", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"repeated single", "aaaa", []string{"a", "aa"}, true},
		{"backtrack prefix", "cars", []string{"car", "ca", "rs"}, true},
		{"no match", "a", []string{"b"}, false},
		{"empty string", "", []string{"a"}, true},
		{"leftover char", "applea", []string{"apple", "pen"}, false},
	}
	for _, c := range cases {
		if got := WordBreak(c.s, c.dict); got != c.want {
			t.Errorf("%s: WordBreak(%q, %v) = %v, want %v", c.name, c.s, c.dict, got, c.want)
		}
	}
}
