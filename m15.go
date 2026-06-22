package main

import (
	"fmt"
)

// func main() {
// 	// 	// fmt.Println("1  *******************************")
// 	// 	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
// 	// 	// fmt.Println("2  *******************************")
// 	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
// 	// 	// fmt.Println("3  *******************************")
// 	// 	// fmt.Println(sortSentence([]int{2, 0, 2, 1, 1, 0}))

// }

func sortSentence(s string) string {
	resArr := []string{}
	word := ""
	for i, l := range s {
		if l == ' ' {
			if l == ' ' && (s[i-1] < '0' || s[i-1] > '9') {
				break
			}
			word = ""
			continue
		}
		if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') {
			word += string(l)
		} else {
			resArr = append(resArr, string(l))
			resArr = append(resArr, word)
		}
	}

	fmt.Println(resArr)
	return sortArr(resArr)
}

func sortArr(arr []string) string {
	res := ""
	for i := 0; i < len(arr)-1; {
		small := i
		fmt.Println(res)
		for k := i + 2; k < len(arr); k += 2 {
			if arr[k] < arr[small] {
				small = k
			}
		}
		res += arr[small+1]
		if len(arr) >= 4 {
			res += " "
		}
		arr = append(arr[:small], arr[small+2:]...)
	}
	fmt.Println("..", res)
	return res
}
