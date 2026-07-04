package main

// import (
// 	"fmt"
// 	"slices"
// )

// // Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func main() {
// 	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))

// }
// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	var output = TreeNode{Val: preorder[0]}
// 	idx := slices.Index(inorder,preorder[0])
// 	preorder= preorder[1:]
// 	// arrRight:= []int{}
// 	arrRight := inorder[idx:]
// 	arrLeft :=  inorder[:idx]
// 	if len(arrRight) != 0 {
// 		buildTree(preorder[])
// 	}

// }
