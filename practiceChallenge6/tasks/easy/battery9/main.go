package main

import "fmt"

// RemoveChar — see README, Task 1. Must be recursive (no loops).
func RemoveChar(s string, c byte) string {
	// TODO: implement
	if len(s) == 0 {
		return s
	}
	if s[0] == c {
		return RemoveChar(s[1:], c)
	}
	return string(s[0]) + RemoveChar(s[1:], c)
}

// ConcatSlice — see README, Task 2. Must be recursive (no loops).
func ConcatSlice(strs []string) string {
	// TODO: implement
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	return strs[0] + ConcatSlice(strs[1:])
}

// Range — see README, Task 3. Must be recursive (no loops).
func Range(start, end int) []int {
	// TODO: implement

	if start > end {
		return []int{}
	}
	if end == start {
		return []int{end}
	}
	r := []int{start}
	res := Range(start-1, end)
	res = append(r, res...)
	return res

}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(RemoveChar("hello", 'h'))
	fmt.Println(ConcatSlice([]string{"a", "b", "c"}))
	fmt.Println(Range(1, 3))
}
