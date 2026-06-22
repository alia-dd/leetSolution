package main

import "fmt"

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("1  *******************************")
// 	fmt.Println(myAtoi("21474836460"))
// 	fmt.Println(myAtoi("-91283472332"))
// 	fmt.Println("2  *******************************")
// 	fmt.Println(myAtoi("  +  413"))
// 	// fmt.Println("3  *******************************")
// 	// fmt.Println(myAtoi("13 37c0d3"))
// 	// fmt.Println("4  *******************************")
// 	// fmt.Println(myAtoi("0-1"))
// 	// fmt.Println("5  *******************************")
// 	// fmt.Println(myAtoi("words and 987"))
// }

func myAtoi(s string) int {

	result := 0
	sing := 1
	found := false
	singfound := false
	for _, n := range s {
		fmt.Println(result)

		if n == ' ' && !found && !singfound {
			continue
		}
		if n == '+' && !found && singfound {
			break
		}
		if n == '+' && !found && !singfound {
			singfound = true
			continue
		}
		if n >= '0' && n <= '9' {
			result = (result * 10) + int(n-'0')
			found = true
			if result >= 2147483648 {
				result = 2147483647
				if sing < 0 {
					result = 2147483647 + 1
				}
				break
			}
			continue
		}
		if n == '-' && !found && !singfound {
			singfound = true
			sing *= -1
			continue
		}
		if n < '0' || n > '9' || n != '+' {
			break
		}
	}

	return result * sing
}
