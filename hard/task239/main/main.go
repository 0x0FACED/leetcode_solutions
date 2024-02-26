package main

import (
	"fmt"
	"leetcode_solutions/hard/task239"
)

func main() {
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	fmt.Println(task239.MaxSlidingWindow(nums1, k1))

	nums2 := []int{1, -1}
	fmt.Println(task239.MaxSlidingWindow_Deque(nums2, 1))
}
