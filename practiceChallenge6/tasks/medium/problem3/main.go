package main

import (
	"fmt"
)

// Flatten — see README. Collect all ints from an arbitrarily-nested []any into
// a flat []int, left-to-right depth-first. Must recurse into nested slices.
func Flatten(nested []any) []int {
	// TODO: implement
	if len(nested) == 0 {
		return []int{}
	}

	var result []int

	switch v := nested[0].(type) {
	case int:
		result = append(result, v)
	case []any:
		result = append(result, Flatten(v)...)

	}
	result = append(result, Flatten(nested[1:])...)
	return result
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Flatten([]any{1, []any{2, 3}, []any{[]any{4}, 5}}))
}
