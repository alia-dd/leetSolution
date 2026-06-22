package main

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := ""
	num2 := ""
	var result *ListNode
	var tail *ListNode
	for l1 != nil {
		fmt.Println(".", l1.Val)
		num1 += strconv.Itoa(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		num2 += strconv.Itoa(l2.Val)
		l2 = l2.Next
	}
	temp := strings.Split(num1, "")
	slices.Reverse(temp)
	num1 = strings.Join(temp, "")
	temp1 := strings.Split(num2, "")
	slices.Reverse(temp1)
	num2 = strings.Join(temp1, "")

	intnum1 := new(big.Int)
	intnum2 := new(big.Int)
	intnum1.SetString(num1, 10)
	intnum2.SetString(num2, 10)
	fmt.Println(intnum1)
	fmt.Println(intnum2)
	intnum1.Add(intnum1, intnum2)

	fmt.Println(intnum1)
	if intnum1.Sign() == 0 {
		result = &ListNode{Val: int(intnum1.Int64())}
	}
	for intnum1.Sign() > 0 {
		mod := new(big.Int)
		intnum1.DivMod(intnum1, big.NewInt(10), mod)
		node := int(mod.Int64())
		newNode := &ListNode{Val: node}
		if result == nil {
			result = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}

	}

	return result
}

// func main(){
// 	var n1 *ListNode
// 	var n2 *ListNode
// 	var tail *ListNode
// 	var tail2 *ListNode
// 	l1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
// 	l2 := []int{5,6,4}
// 	for _, n := range l1{
// 		newNode := &ListNode{Val : n}
// 		if n1 == nil{
// 			n1 = newNode
// 			tail = newNode
// 		}else{
// 			tail.Next =  newNode
// 			tail = newNode
// 		}
// 	}
// 	for _, n := range l2{
// 		newNode := &ListNode{Val : n}
// 		if n2 == nil{
// 			n2 = newNode
// 			tail2 = newNode
// 		}else{
// 			tail2.Next =  newNode
// 			tail2 = newNode
// 		}
// 	}
// 	fmt.Println(addTwoNumbers(n1,n2))
// }
