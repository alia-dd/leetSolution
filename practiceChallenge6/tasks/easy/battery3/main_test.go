package main

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{7, 13},
		{10, 55},
	}
	for _, c := range cases {
		if got := Fibonacci(c.n); got != c.want {
			t.Errorf("Fibonacci(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}

func TestGCD(t *testing.T) {
	cases := []struct {
		a, b int
		want int
	}{
		{12, 8, 4},
		{48, 36, 12},
		{7, 13, 1},
		{10, 0, 10},
		{0, 5, 5},
		{100, 25, 25},
	}
	for _, c := range cases {
		if got := GCD(c.a, c.b); got != c.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestCountDown(t *testing.T) {
	cases := []struct {
		n    int
		want []int
	}{
		{0, []int{0}},
		{1, []int{1, 0}},
		{3, []int{3, 2, 1, 0}},
		{5, []int{5, 4, 3, 2, 1, 0}},
	}
	for _, c := range cases {
		if got := CountDown(c.n); !reflect.DeepEqual(got, c.want) {
			t.Errorf("CountDown(%d) = %v, want %v", c.n, got, c.want)
		}
	}
}
