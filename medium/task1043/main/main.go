package main

import (
	"fmt"
	"leetcode_solutions/medium/task1043"
)

func main() {
	arr1 := []int{1, 15, 7, 9, 2, 5, 10}
	k1 := 3
	fmt.Println(task1043.MaxSumAfterPartitioning(arr1, k1))

	arr2 := []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}
	k2 := 4
	fmt.Println(task1043.MaxSumAfterPartitioning(arr2, k2))

	arr3 := []int{1}
	k3 := 1
	fmt.Println(task1043.MaxSumAfterPartitioning(arr3, k3))
}
