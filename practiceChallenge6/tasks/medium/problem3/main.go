package main

import "fmt"

// Flatten — see README. Collect all ints from an arbitrarily-nested []any into
// a flat []int, left-to-right depth-first. Must recurse into nested slices.
func Flatten(nested []any) []int {
	// TODO: implement
	return nil
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Flatten([]any{1, []any{2, 3}, []any{[]any{4}, 5}}))
}
