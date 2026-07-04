package main

// func main() {
// 	fmt.Println(convert("PAYPALISHIRING", 4))
// 	// fmt.Println(convert("PAYPALISHIRING", 3))

// }

// func convert(s string, numRows int) string {
// 	res := ""
// 	if numRows == 1 {
// 		return s
// 	}
// 	for i := 0; i < len(s); i++ {
// 		j := numRows
// 		row := 1
// 		if numRows > 2 {
// 			j = numRows + (numRows - 2)
// 		}
// 		res += string(s[i])
// 		if row != 1 || row != numRows {
// 			j++
// 		}
// 		fmt.Println(j)

// 		for {

// 			if i+j > len(s)-1 {
// 				break
// 			}

// 			fmt.Println(string(s[i+j]))
// 			res += string(s[i+j])
// 			j += j

// 		}
// 		row++
// 		fmt.Println(res)
// 		res = ""
// 	}
// 	return res
// }
