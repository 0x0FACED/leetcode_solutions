package main

import (
	"fmt"
	"leetcode_solutions/easy/task100"
	"leetcode_solutions/internal/treenode"
)

func main() {
	arr1 := []interface{}{1, 2, 3, 5, 6, 7, 8, nil, nil, 10}
	arr2 := []interface{}{1, 2, 3, 5, 6, 2, 8, nil, nil, 10}
	p := treenode.FillTree(arr1)
	q := treenode.FillTree(arr2)
	fmt.Println(task100.IsSameTree(p, q))

}
