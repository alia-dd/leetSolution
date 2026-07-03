package main

import "fmt"

// Factorial — see README, Task 1. Must be recursive (no loops).
func Factorial(n int) int {
	// TODO: implement
	if n < 1 {
		return 1
	}
	return Factorial(n-1) * n
}

// SumDigits — see README, Task 2. Must be recursive (no loops).
func SumDigits(n int) int {
	// TODO: implement
	if n < 0 {
		return -n
	}
	if n < 10 {
		return n
	}
	return n%10 + SumDigits(n/10)
}

// Power — see README, Task 3. Must be recursive (no loops).
func Power(base, exp int) int {
	// TODO: implement
	if exp == 0 {
		return 1
	}
	return Power(base, exp-1) * base
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Factorial(5))
	fmt.Println(SumDigits(123))
	fmt.Println(Power(2, 10))
}
