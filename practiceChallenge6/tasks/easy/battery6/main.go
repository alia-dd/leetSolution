package main

import "fmt"

// Multiply — see README, Task 1. Repeated addition, no `*`. Must be recursive.
func Multiply(a, b int) int {
	// TODO: implement
	if b <= 0 {
		return 0
	}
	if b <= 1 {
		return a
	}
	return a + Multiply(a, b-1)
}

// CountDigits — see README, Task 2. Must be recursive (no loops).
func CountDigits(n int) int {
	// TODO: implement
	if n < 10 {
		return 1
	}
	return 1 + CountDigits(n/10)
}

// SumTo — see README, Task 3. Must be recursive (no loops).
func SumTo(n int) int {
	// TODO: implement
	if n <= 1 {
		return n
	}
	return n + SumTo(n-1)
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Multiply(3, 4))
	fmt.Println(CountDigits(12345))
	fmt.Println(SumTo(5))
}
