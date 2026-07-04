package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// fmt.Println(*buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	fmt.Print(*buildTree([]int{4, 5, 3, 17, 2, 6, 7, 1, 9, 8, 10, 11, 13, 14, 15, 20, 19}, []int{3, 2, 6, 17, 1, 7, 5, 9, 8, 4, 10, 14, 13, 11, 20, 15, 19}))
	// buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}).String()

}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}
	// if len(preorder) == 1 {
	// 	return &TreeNode{Val: 9}
	// }
	t := &TreeNode{Val: preorder[0]}
	index := indexOf(t.Val, inorder)
	leftIn := inorder[:index]
	rightin := inorder[index+1:]
	leftPRe := preorder[1 : len(leftIn)+1]
	righPRe := preorder[len(leftIn)+1:]

	t.Left = buildTree(leftPRe, leftIn)
	t.Right = buildTree(righPRe, rightin)

	// fmt.Println(leftIn, rightin)
	// fmt.Println(leftPRe, righPRe)
	return t
}

func (node *TreeNode) String() string {
	if node == nil {
		return "NULL"
	}
	return fmt.Sprintf("%d Left:{%s} Right:{%s}", node.Val, node.Left.String(), node.Right.String())
}

func indexOf(element int, data []int) int {
	// it doesnt handle same value in multiple postion will take the first one
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
