package main

import (
	"reflect"
	"testing"
)

func TestHanoiSmall(t *testing.T) {
	cases := []struct {
		n    int
		want []string
	}{
		{1, []string{"A->C"}},
		{2, []string{"A->B", "A->C", "B->C"}},
		{3, []string{"A->C", "A->B", "C->B", "A->C", "B->A", "B->C", "A->C"}},
	}
	for _, c := range cases {
		if got := Hanoi(c.n, "A", "C", "B"); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Hanoi(%d, A, C, B) =\n %v\nwant\n %v", c.n, got, c.want)
		}
	}
}

func TestHanoiZero(t *testing.T) {
	if got := Hanoi(0, "A", "C", "B"); len(got) != 0 {
		t.Errorf("Hanoi(0, ...) = %v, want no moves", got)
	}
}

func TestHanoiMoveCount(t *testing.T) {
	// The move count must be exactly 2^n - 1 for any n.
	for n := 0; n <= 10; n++ {
		want := (1 << n) - 1
		if got := len(Hanoi(n, "A", "C", "B")); got != want {
			t.Errorf("Hanoi(%d) produced %d moves, want %d", n, got, want)
		}
	}
}

func TestHanoiCustomPegNames(t *testing.T) {
	// Peg names are arbitrary; the structure must follow from/to/via.
	want := []string{"X->Z", "X->Y", "Z->Y", "X->Z", "Y->X", "Y->Z", "X->Z"}
	if got := Hanoi(3, "X", "Z", "Y"); !reflect.DeepEqual(got, want) {
		t.Errorf("Hanoi(3, X, Z, Y) =\n %v\nwant\n %v", got, want)
	}
}
