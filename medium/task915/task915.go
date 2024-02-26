package main

import (
	"fmt"
	"math"
)

func partitionDisjoint(nums []int) int {

	length := len(nums)
	maxLeft := math.MinInt
	maxRight := math.MinInt
	right := length - 1
	for i := 0; i < right; i++ {
		if nums[i] > maxLeft {
			maxLeft = nums[i]
		}
		if nums[right] > maxRight {
			maxRight = nums[i]
		}
		if maxRight < maxLeft {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
}
