package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	// fmt.Println(FormatReceipt("Coffee", 3.8))
// 	// fmt.Println(FormatReceipt("Sandwich", 12.999))
// 	// fmt.Println("**************************************")
// 	// fmt.Println(FormatTable([]string{"Coffee", "pagel", "Latte", "Sandwich"}, []float64{3.5, 2.99, 9.500, 12.9999}))
// 	// fmt.Println("**************************************")
// 	// fmt.Println(NormalizeAmount("895,900,999,100,234.56"))
// 	// fmt.Println(NormalizeAmount("995,599,404,339,234"))
// 	// fmt.Println(NormalizeAmount("895,900,99,100,234.56"))
// 	fmt.Println(NormalizeAmount("1,234.56"))     // -> "1234.56", true
// 	fmt.Println(NormalizeAmount("1,000,000.99")) // -> "1000000.99", true
// 	fmt.Println(NormalizeAmount("42.5"))         // -> "42.50", true
// 	fmt.Println(NormalizeAmount("0.99"))         // -> "0.99", true
// 	fmt.Println("**************************************")
// 	// Invalid cases
// 	fmt.Println(NormalizeAmount("1,23.56")) // -> "", false  (wrong grouping, 2 digits before comma)
// 	fmt.Println(NormalizeAmount("abc"))     // -> "", false
// 	fmt.Println(NormalizeAmount("1.2.3"))   // -> "", false  (two decimal points)
// 	fmt.Println(NormalizeAmount(""))        // -> "", false
// 	fmt.Println(NormalizeAmount("1,234"))   // -> "1234.00", true  (no decimal, still valid)

// }

// func FormatReceipt(item string, price float64) string {
// 	// solution
// 	return fmt.Sprintf("%-20s$%.2f ", item, price)
// }

func FormatTable(items []string, prices []float64) string {
	// solution
	var totp float64
	for i, p := range prices {
		line := fmt.Sprintf("%-20s %.2f", items[i], p)
		totp += p
		fmt.Println(line)
	}
	return fmt.Sprintf("Total: %23.2f", totp)
}

// 1c — Strict format parsing/validation, no fmt package for parsing
// Allowed packages: none for parsing logic (fmt allowed only for output).

// Write a function that takes a custom numeric string format
// (e.g. "1,234.56" with thousands separators) and converts it
// to a clean float64, then re-formats it back with padding rules.

func NormalizeAmount(input string) (string, bool) {
	// solution
	var res strings.Builder
	var num strings.Builder
	lnum := ""
	count := 0
	after := false
	for i, n := range input {
		if (n == ',' && (i == 0)) || (n == '.' && after) {
			fmt.Println(num.String(), i, count)
			return "Invalid Input.", false
		}
		if n == '.' {
			after = true
			continue
		}
		if !after {
			if n == ',' {
				num.WriteString(" ")
			} else {
				num.WriteString(string(n))
			}
		} else {
			lnum += string(n)
		}
		if count == 3 {
			count = 0
			continue
		}
		count++
	}
	res.WriteString(num.String())
	res.WriteString(".")
	// fmt.Println(len(lnum))
	if len(lnum) > 1 {
		res.WriteString(lnum[:2])
	} else {
		res.WriteString(lnum)
		for range 2 - len(lnum) {
			res.WriteString("0")
		}
	}
	return res.String(), true
}
