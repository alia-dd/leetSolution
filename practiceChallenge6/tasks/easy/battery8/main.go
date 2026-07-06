package main

import (
	"fmt"
)

// IsEven — see README, Task 1. No % or /. Must be recursive.
func IsEven(n int) bool {
	// TODO: implement
	switch n {
	case 0:
		return true
	case 1:
		return false
	}

	return IsEven(n - 2)
}

// IsPowerOfTwo — see README, Task 2. Must be recursive (no loops).
func IsPowerOfTwo(n int) bool {
	// TODO: implement
	if n == 1 {
		return true
	}
	if n%2 != 0 {
		return false
	}
	return IsPowerOfTwo(n / 2)
}

// CollatzSteps — see README, Task 3. Must be recursive (no loops).
// i dont get this a bit difficult to undestand
func CollatzSteps(n int) int {
	// TODO: implement
	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + CollatzSteps(n/2)
	}
	return 1 + CollatzSteps(3*n+1)
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(IsEven(3))
	fmt.Println(IsPowerOfTwo(3))
	fmt.Println(CollatzSteps(6))

}
