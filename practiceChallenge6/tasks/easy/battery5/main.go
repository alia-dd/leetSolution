package main

import "fmt"

// Repeat — see README, Task 1. Must be recursive (no loops).
func Repeat(s string, n int) string {
	// TODO: implement
	return ""
}

// CountChar — see README, Task 2. Must be recursive (no loops).
func CountChar(s string, target byte) int {
	// TODO: implement
	return 0
}

// IsSorted — see README, Task 3. Must be recursive (no loops).
func IsSorted(nums []int) bool {
	// TODO: implement
	return false
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Repeat("ab", 3))
	fmt.Println(CountChar("hello", 'l'))
	fmt.Println(IsSorted([]int{1, 2, 2, 3}))
}
