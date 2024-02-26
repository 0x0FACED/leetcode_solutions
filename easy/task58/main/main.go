package main

import (
	"fmt"
	"leetcode_solutions/easy/task58"
)

func main() {
	// Task 58
	//str1 := "   fly me   to   the moon  "
	str2 := "luffy is still joyboy"
	fmt.Println(task58.LengthOfLastWord(str2))
	fmt.Println(task58.LengthOfLastWord_Fields(str2))
}
