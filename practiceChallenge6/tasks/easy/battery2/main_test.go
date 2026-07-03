package main

import "testing"

func TestCountVowels(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"xyz", 0},
		{"hello", 2},
		{"aeiou", 5},
		{"AEIOU", 5},
		{"Recursion", 4},
	}
	for _, c := range cases {
		if got := CountVowels(c.in); got != c.want {
			t.Errorf("CountVowels(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"a", true},
		{"abba", true},
		{"racecar", true},
		{"abcba", true},
		{"hello", false},
		{"ab", false},
		{"Anna", false}, // case-sensitive
	}
	for _, c := range cases {
		if got := IsPalindrome(c.in); got != c.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestSumSlice(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{1, 2, 3}, 6},
		{[]int{-1, 1}, 0},
		{[]int{10, 20, 30, 40}, 100},
	}
	for _, c := range cases {
		if got := SumSlice(c.in); got != c.want {
			t.Errorf("SumSlice(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
