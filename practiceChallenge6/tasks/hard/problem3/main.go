package main

import "fmt"

// WordBreak — see README. Report whether s can be segmented into a sequence of
// dictionary words (reusable, whole string consumed). Must be recursive.
func WordBreak(s string, dict []string) bool {
	// TODO: implement
	return false
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(WordBreak("applepenapple", []string{"apple", "pen"}))
}
