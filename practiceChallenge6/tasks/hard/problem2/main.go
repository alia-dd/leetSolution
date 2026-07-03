package main

import "fmt"

// Evaluate — see README. Parse and evaluate an integer arithmetic expression
// (+, -, *, /, parentheses, precedence). Return (value, true) or (0, false) if
// malformed. Implement with recursive descent.
func Evaluate(expr string) (int, bool) {
	// TODO: implement
	return 0, false
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Evaluate("2*(3+4)-5"))
}
