package main

import "fmt"

// Hanoi — see README. Return the ordered list of "X->Y" moves that transfers n
// disks from `from` to `to` using `via` as the spare. Must be recursive.
func Hanoi(n int, from, to, via string) []string {
	// TODO: implement
	return nil
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Hanoi(3, "A", "C", "B"))
}
