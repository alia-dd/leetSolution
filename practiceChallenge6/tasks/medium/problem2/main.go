package main

import "fmt"

// Permutations — see README. Return all n! orderings of the characters of s
// (duplicates kept). Must be recursive.
func Permutations(s string) []string {
	// TODO: implement
	if len(s) <= 1 {
		return []string{s}
	}

	f := string(s[:1])
	res := Permutations(s[1:])

	// res1 := Permutations(s[2:])
	fmt.Println(f, ">", res)
	// Permutations()
	return res
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Permutations("abcd"))
}
