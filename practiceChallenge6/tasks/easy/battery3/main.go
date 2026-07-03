package main

import "fmt"

// Fibonacci — see README, Task 1. Must be recursive (no loops).
func Fibonacci(n int) int {
	// TODO: implement
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// GCD — see README, Task 2. Must be recursive (no loops).
func GCD(a, b int) int {
	// Step 1: \(48 \div 18 = 2\) with a remainder of \(12\) (\(48 = 18 \cdot 2 + 12\))
	// Step 2: \(18 \div 12 = 1\) with a remainder of \(6\) (\(18 = 12 \cdot 1 + 6\))
	// Step 3: \(12 \div 6 = 2\) with a remainder of \(0\) (\(12 = 6 \cdot 2 + 0\))
	// TODO: implement
	// if a%b == 0 {
	// 	return b
	// }
	// res := b*a/b + (a % b)
	return 0
}

// CountDown — see README, Task 3. Must be recursive (no loops).
func CountDown(n int) []int {
	// TODO: implement
	if n == 0 {
		return []int{0}
	}
	res := CountDown(n - 1)
	r := []int{n}
	res = append(r, res...)
	return res
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Fibonacci(10))
	fmt.Println(GCD(48, 36))
	fmt.Println(CountDown(3))
}
