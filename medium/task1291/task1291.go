package task1291

import (
	"sort"
)

// Example 1:
// Input: low = 100, high = 300
// Output: [123,234]
//
// Example 2:
// Input: low = 1000, high = 13000
// Output: [1234,2345,3456,4567,5678,6789,12345]
func SequentialDigits(low int, high int) []int {
	var result []int

	for i := 1; i <= 9; i++ {
		num := i
		for j := i + 1; j <= 9; j++ {
			num = num*10 + j
			if num >= low && num <= high {
				result = append(result, num)
			}
		}
	}
	sort.Ints(result)
	return result
}
