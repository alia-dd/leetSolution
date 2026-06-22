package main

import (
	"fmt"
	"strconv"
	"strings"
)

//	func main() {
//		fmt.Println(convertDateToBinary("2080-02-29"))
//		fmt.Println("*******************************")
//		fmt.Println(convertDateToBinary("900-01-01"))
//		fmt.Println("*******************************")
//		// fmt.Println(convertDateToBinary("26th May 1960"))
//	}
func convertDateToBinary(date string) string {
	if len(date) != 10 || date[4] != '-' || date[7] != '-' {
		return ""
	}

	list := strings.Split(date, "-")
	for i, t := range list {
		list[i] = tobinary(t)
	}
	res := strings.Join(list, "-")
	return res
}
func tobinary(data string) string {
	bi := ""
	num, _ := strconv.Atoi(data)
	for num > 0 {
		bi = strconv.Itoa(num%2) + bi
		num = num / 2
		fmt.Println(num)
	}
	fmt.Println("bu", bi)
	return bi
}
