package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// func main() {

// 	// fmt.Println(sort(654321))
// 	fmt.Println(Itoa(654321))
// 	fmt.Println(Atoi("654321"))
// }

func sort(num int) int {
	fmt.Println(num)
	st := strconv.Itoa(num)
	s := strings.Split(st, "")
	slices.Sort(s)
	fmt.Println(s)
	st = strings.Join(s, "")
	n, _ := strconv.Atoi(st)
	return n
}

func Atoi(str string) int {
	num := 0
	for _, s := range str {
		num *= 10
		num += int(s - '0')

	}
	return num
}

func Itoa(num int) string {
	str := ""
	if num == 0 {
		return string(rune(num + '0'))
	}
	for ; num > 0; num /= 10 {
		digit := num % 10
		str = string(rune(digit+'0')) + str
	}
	return str

}
