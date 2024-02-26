package treenode

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewEmpty() *TreeNode {
	return &TreeNode{}
}

func NewWithVal(x int) *TreeNode {
	return &TreeNode{
		Val: x,
	}
}

func FillTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	var root *TreeNode
	root = fillTreeHelper(arr, root, 0)
	return root
}

func fillTreeHelper(arr []interface{}, root *TreeNode, i int) *TreeNode {
	if i < len(arr) {
		if arr[i] == nil {
			return nil
		}
		temp := &TreeNode{Val: arr[i].(int)}
		root = temp

		root.Left = fillTreeHelper(arr, root.Left, 2*i+1)
		root.Right = fillTreeHelper(arr, root.Right, 2*i+2)
	}
	return root
}

func printInOrder(root *TreeNode) {
	if root != nil {
		printInOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		printInOrder(root.Right)
	}
}
