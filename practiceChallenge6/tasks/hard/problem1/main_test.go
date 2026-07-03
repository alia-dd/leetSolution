package main

import "testing"

func TestCountNQueens(t *testing.T) {
	// index = n, value = number of solutions (n from 1 to 8)
	want := map[int]int{
		1: 1,
		2: 0,
		3: 0,
		4: 2,
		5: 10,
		6: 4,
		7: 40,
		8: 92,
	}
	for n, w := range want {
		if got := CountNQueens(n); got != w {
			t.Errorf("CountNQueens(%d) = %d, want %d", n, got, w)
		}
	}
}
