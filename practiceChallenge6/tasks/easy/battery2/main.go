package main

import "fmt"

// CountVowels — see README, Task 1. Must be recursive (no loops).
func CountVowels(s string) int {
	// TODO: implement
	if len(s) < 1 {
		return 0
	}
	if len(s) == 1 {
		if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" || s == "A" || s == "E" || s == "I" || s == "O" || s == "U" {
			return 1
		} else {
			return 0
		}
	}
	n := 0
	if s[:1] == "a" || s[:1] == "e" || s[:1] == "i" || s[:1] == "o" || s[:1] == "u" || s[:1] == "A" || s[:1] == "E" || s[:1] == "I" || s[:1] == "O" || s[:1] == "U" {
		n = 1
	}
	return n + CountVowels(s[1:])
}

// IsPalindrome — see README, Task 2. Must be recursive (no loops).
func IsPalindrome(s string) bool {
	// TODO: implement
	if len(s) <= 1 {
		return true
	}
	if len(s) <= 3 {
		if s[0] == s[len(s)-1] {
			return true
		}
	}

	if IsPalindrome(s[1:len(s)-1]) && s[0] == s[len(s)-1] {
		return true
	}
	return false
}

// SumSlice — see README, Task 3. Must be recursive (no loops).
func SumSlice(nums []int) int {
	// TODO: implement
	if len(nums) < 1 {
		return 0
	}
	return nums[0] + SumSlice(nums[1:])
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(CountVowels("Recursion"))
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("hello"))
	fmt.Println(SumSlice([]int{1, 2, 3}))
}
