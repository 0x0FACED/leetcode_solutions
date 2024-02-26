package main

import (
	"fmt"
	"leetcode_solutions/medium/task1291"
)

func main() {
	// Примеры использования:
	low1, high1 := 100, 300
	fmt.Println(task1291.SequentialDigits(low1, high1))

	low2, high2 := 1000, 130000000
	fmt.Println(task1291.SequentialDigits(low2, high2))
}
