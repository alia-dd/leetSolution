package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println(convert("PAYPALISHIRING", 4))
// 	// Output: "PINALSIGYAHRPI"
// 	fmt.Println(convert("PAYPALISHIRING", 3))
// 	// Output: "PAHNAPLSIIGYIR"
// 	fmt.Println(convert("PAYPALISHIRING", 1))
// 	// Output: "PAYPALISHIRING"

// }

func convert(s string, numRows int) string {
	resArr := [10][10]string{}
	res := ""
	dio := false
	if numRows == 1 {
		return s
	}
	col, row := 0, 0
	for i := 0; i < len(s); i++ {
		if row == numRows {
			row = numRows - 2
			col++
			dio = true
		}
		// for diagonal handling
		if dio {
			if row == 0 {
				resArr[row][col] = string(s[i])
				row++
				dio = false
				continue
			}
			resArr[row][col] = string(s[i])
			row--
			col++
		} else {
			resArr[row][col] = string(s[i])
			row++
		}

	}
	// add to res for output
	for i, arr := range resArr {
		for k, v := range arr {
			res += v
			if v == "" {
				resArr[i][k] = "#"
			}
		}
		fmt.Println(arr)
	}

	return res
}
