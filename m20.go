package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	fmt.Println("res:", numDecodings("226")) /// 2 2 6,  22 6 , 2 26
// 	fmt.Println("-------------------------------")
// 	fmt.Println("res:", numDecodings("11106"))
// }

func numDecodings(s string) int {
	counti := 0
	for i := 0; i < len(s); i++ {

		n, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return 0
		}
		fmt.Println(n)
		// fmt.Println(count)
		if n < 0 {
			continue
		}
		counti++

		if i+1 < len(s) {
			n2, err := strconv.Atoi(string(s[i+1]))
			if err != nil || n2 < 0 {
				continue
			}
			if (n*10)+n2 < 27 {
				// fmt.Println("(n*10)+n2 < 27", (n*10)+n2)
				counti++

			}
		}
	}
	return counti
}
