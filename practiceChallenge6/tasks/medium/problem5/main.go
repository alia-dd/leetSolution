package main

import (
	"fmt"
)

// BinarySearch — see README. Return the index of target in the ascending,
// duplicate-free slice, or -1 if absent. Implement recursively.
func BinarySearch(sorted []int, target int) int {
	// TODO: implement
	if len(sorted) == 0 {
		return -1
	}
	mid := len(sorted) / 2

	if sorted[mid] == target {
		return mid
	}
	if target < sorted[mid] {
		return BinarySearch(sorted[:mid], target)
	}
	res := BinarySearch(sorted[mid+1:], target)
	if res == -1 {
		return res
	}

	return mid + res + 1
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(BinarySearch([]int{1, 3, 5, 7, 9}, 7))
}
