package main

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{5, 120},
		{10, 3628800},
	}
	for _, c := range cases {
		if got := Factorial(c.n); got != c.want {
			t.Errorf("Factorial(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}

func TestSumDigits(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{0, 0},
		{7, 7},
		{10, 1},
		{123, 6},
		{99999, 45},
		{1000000, 1},
	}
	for _, c := range cases {
		if got := SumDigits(c.n); got != c.want {
			t.Errorf("SumDigits(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}

func TestPower(t *testing.T) {
	cases := []struct {
		base, exp int
		want      int
	}{
		{2, 0, 1},
		{0, 0, 1},
		{2, 10, 1024},
		{5, 3, 125},
		{3, 1, 3},
		{-2, 3, -8},
		{-2, 2, 4},
	}
	for _, c := range cases {
		if got := Power(c.base, c.exp); got != c.want {
			t.Errorf("Power(%d, %d) = %d, want %d", c.base, c.exp, got, c.want)
		}
	}
}
