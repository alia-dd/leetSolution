package main

import (
	"reflect"
	"sort"
	"testing"
)

// sortedEqual compares two string slices ignoring order.
func sortedEqual(a, b []string) bool {
	a = append([]string(nil), a...)
	b = append([]string(nil), b...)
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}

func TestPermutations(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"", []string{""}},
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		// duplicates are kept: "aab" -> 6 entries, not 3
		{"aab", []string{"aab", "aab", "aba", "aba", "baa", "baa"}},
	}
	for _, c := range cases {
		if got := Permutations(c.in); !sortedEqual(got, c.want) {
			t.Errorf("Permutations(%q) = %v, want (any order) %v", c.in, got, c.want)
		}
	}
}

func TestPermutationsCount(t *testing.T) {
	// length must be n! for a length-n input.
	fact := []int{1, 1, 2, 6, 24, 120} // 0!..5!
	for n := 0; n <= 5; n++ {
		s := "abcde"[:n]
		if got := len(Permutations(s)); got != fact[n] {
			t.Errorf("len(Permutations(%q)) = %d, want %d", s, got, fact[n])
		}
	}
}
