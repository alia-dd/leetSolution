package main

import "fmt"

// Hanoi — see README. Return the ordered list of "X->Y" moves that transfers n
// disks from `from` to `to` using `via` as the spare. Must be recursive.
func Hanoi(n int, from, to, via string) []string {
	// TODO: implement
	if n <= 0 {
		return []string{}
	}

	move := []string{fmt.Sprintf("%s->%s", from, via)}
	move2 := []string{fmt.Sprintf("%s->%s", from, to)}

	move3 := []string{fmt.Sprintf("%s->%s", via, to)}
	move = append(move, move2...)

	move = append(move3, Hanoi(n-1, from, to, via)...)
	return move
}

// dont get it

// func Hanoi(n int, from, to, via string) []string {
// 	if n <= 0 {
// 		return []string{}
// 	}

// 	// 1. move n-1 disks from "from" to "via"
// 	left := Hanoi(n-1, from, via, to)

// 	// 2. move the biggest disk from "from" to "to"
// 	move := []string{from + "->" + to}

// 	// 3. move n-1 disks from "via" to "to"
// 	right := Hanoi(n-1, via, to, from)

//		return append(append(left, move...), right...)
//	}
func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Hanoi(3, "A", "C", "B"))
}
