package main

import "fmt"

// Permutations — see README. Return all n! orderings of the characters of s
// (duplicates kept). Must be recursive.
func Permutations(s string) []string {
	// TODO: implement
	if len(s) <= 1 {
		return []string{s}
	}

	f := string(s[0])
	res := Permutations(s[1:])

	var result []string

	for _, p := range res {
		for i := 0; i <= len(p); i++ {
			newStr := p[:i] + f + p[i:]
			result = append(result, newStr)
		}
	}
	return result
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Permutations("abc"))
}
