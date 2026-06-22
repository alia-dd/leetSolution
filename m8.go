package main

import (
	"fmt"
	"strings"
	"time"
)

//	func main() {
//		fmt.Println(reformatDate("20th Oct 2052"))
//		fmt.Println("*******************************")
//		fmt.Println(reformatDate("6th Jun 1933"))
//		fmt.Println("*******************************")
//		fmt.Println(reformatDate("26th May 1960"))
//	}
func reformatDate(date string) string {
	// pattern := regexp.MustCompile("\\d{1,2}(?:st|nd|rd|th)\\s([A-Za-z]){3}\\s\\d{4}")
	// result := pattern.FindAllString(date, -1)
	// if len(result) != 3 {
	// 	return ""
	// // }
	// fmt.Println(result[0])
	// fmt.Println(len(result))
	// year, err := strconv.Atoi(result[2])
	// if err != nil || year < 1900 || year > 2100 {
	// 	return ""
	// }
	clean := strings.Replace(date, "th", "", 1)
	clean = strings.Replace(clean, "nd", "", 1)
	clean = strings.Replace(clean, "rd", "", 1)
	fmt.Println(clean)
	dateres, _ := time.Parse("2 Jan 2006", clean)
	dateres.Format("2006-01-02")
	return dateres.String()
}
