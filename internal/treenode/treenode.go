package treenode

import "fmt"

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

// Метод для заполнения дерева из входного массива
func FillTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	var root *TreeNode
	root = fillTreeHelper(arr, root, 0)
	return root
}

func fillTreeHelper(arr []int, root *TreeNode, i int) *TreeNode {
	if i < len(arr) {
		temp := &TreeNode{Val: arr[i]}
		root = temp

		root.Left = fillTreeHelper(arr, root.Left, 2*i+1)
		root.Right = fillTreeHelper(arr, root.Right, 2*i+2)
	}
	return root
}

// Вспомогательная функция для обхода дерева и вывода значений
func printInOrder(root *TreeNode) {
	if root != nil {
		printInOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		printInOrder(root.Right)
	}
}
