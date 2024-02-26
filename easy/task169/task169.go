package task169

import (
	"sort"
)

func MajorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	currEl := nums[0]
	prevEl := 0
	currLength := 1
	prevLength := currLength
	for _, el := range nums[1:] {
		if el != currEl {
			if currLength > prevLength {
				prevLength = currLength
				prevEl = currEl
			}
			currEl = el
			currLength = 1
		} else {
			currLength++
		}

	}

	if currLength > prevLength {
		return currEl
	}
	return prevEl
}
