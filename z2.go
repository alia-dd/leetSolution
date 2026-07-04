package main

// func convert(s string, numRows int) string {
// 	var vertical []string
// 	var diagonal []string
// 	var twod [10][10]string
// 	skip := numRows - 2
// 	col := 0
// 	for i := 0; i < len(s); i += skip {

// 		l := len(s[i:len(s)])
// 		// fmt.Println(l)
// 		if l < numRows {
// 			skip = l
// 			vertical = append(vertical, s[i:i+skip])
// 			break
// 		}
// 		// fmt.Println(s[i : i+numRows])
// 		vertical = append(vertical, s[i:i+numRows])
// 		for i, v := range vertical {
// 			twod[i][0] = v
// 		}
// 		fmt.Println(twod)
// 		i += numRows
// 		if i == len(s) {
// 			break
// 		}
// 		diagonal = append(diagonal, s[i:i+skip])
// 		col++
// 		for i, d := range diagonal {

// 			r := numRows - i - 2
// 			twod[r][col] = d
// 			col++

// 		}

// 	}
// 	// fmt.Println(vertical)
// 	// fmt.Println(diagonal)
// 	for _, t := range twod {
// 		fmt.Println(t)
// 	}
// 	return ""
// }
// func main() {
// 	s := "PAYPALISHIRING"
// 	fmt.Println(convert(s, 4))
// }
