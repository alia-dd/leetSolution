package main

import "testing"

func TestEvaluateValid(t *testing.T) {
	cases := []struct {
		expr string
		want int
	}{
		{"1+2", 3},
		{"2+3*4", 14},
		{"(2+3)*4", 20},
		{"2*(3+4)-5", 9},
		{"10/3", 3},
		{"20-5-3", 12},
		{"8/2/2", 2},
		{"2*3+4*5", 26},
		{"3-5", -2},
		{" 3 + 4 ", 7},
		{"100", 100},
		{"((1+2)*(3+4))", 21},
	}
	for _, c := range cases {
		got, ok := Evaluate(c.expr)
		if !ok || got != c.want {
			t.Errorf("Evaluate(%q) = %d, %v; want %d, true", c.expr, got, ok, c.want)
		}
	}
}

func TestEvaluateInvalid(t *testing.T) {
	bad := []string{
		"",
		"   ",
		"1+",
		"*3",
		"1 2",
		"()",
		"(1+2",
		"1+2)",
		"1++2",
		"a+1",
		"1/0",
		"-3",
	}
	for _, expr := range bad {
		if got, ok := Evaluate(expr); ok || got != 0 {
			t.Errorf("Evaluate(%q) = %d, %v; want 0, false", expr, got, ok)
		}
	}
}
