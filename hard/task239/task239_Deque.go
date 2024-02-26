package task239

import (
	"container/list"
	"fmt"
)

func MaxSlidingWindow_Deque(nums []int, k int) []int {
	result := make([]int, 0)
	maxDeque := list.New()

	for i := 0; i < len(nums); i++ {
		for maxDeque.Len() > 0 && maxDeque.Front().Value.(int) < i-k+1 {
			maxDeque.Remove(maxDeque.Front())
		}

		for maxDeque.Len() > 0 && nums[maxDeque.Back().Value.(int)] < nums[i] {
			maxDeque.Remove(maxDeque.Back())
		}

		maxDeque.PushBack(i)

		if i >= k-1 {
			printDeque(maxDeque)
			result = append(result, nums[maxDeque.Front().Value.(int)])
		}
	}

	return result
}

func printDeque(deque *list.List) {
	for element := deque.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value, " ")
	}
	fmt.Println()
}
