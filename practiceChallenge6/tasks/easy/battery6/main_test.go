package main

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		a, b int
		want int
	}{
		{3, 4, 12},
		{5, 0, 0},
		{0, 7, 0},
		{-2, 3, -6},
		{7, 1, 7},
		{6, 6, 36},
	}
	for _, c := range cases {
		if got := Multiply(c.a, c.b); got != c.want {
			t.Errorf("Multiply(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestCountDigits(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{0, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{12345, 5},
	}
	for _, c := range cases {
		if got := CountDigits(c.n); got != c.want {
			t.Errorf("CountDigits(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}

func TestSumTo(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 15},
		{10, 55},
		{100, 5050},
	}
	for _, c := range cases {
		if got := SumTo(c.n); got != c.want {
			t.Errorf("SumTo(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}
