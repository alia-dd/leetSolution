package main

import "fmt"

// GenerateParentheses — see README. Return all well-formed strings of n pairs of
// parentheses (any order). Must be recursive (backtracking).
func GenerateParentheses(n int) []string {
	// TODO: implement
	if n == 0 {
		return []string{}
	}
	// var result []string{}

	return nil
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(GenerateParentheses(3))
}
