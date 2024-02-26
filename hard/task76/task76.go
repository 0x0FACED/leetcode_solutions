package task76

import "math"

func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	charCount := make(map[byte]int)
	for i := range t {
		charCount[t[i]]++
	}

	requiredChars := len(charCount)
	formedChars := 0

	left, right := 0, 0
	minLength := math.MaxInt64
	minWindow := ""

	for right < len(s) {
		char := s[right]
		charCount[char]--

		if charCount[char] == 0 {
			formedChars++
		}

		for formedChars == requiredChars {
			if right-left+1 < minLength {
				minLength = right - left + 1
				minWindow = s[left : right+1]
			}

			char = s[left]
			charCount[char]++

			if charCount[char] > 0 {
				formedChars--
			}

			left++
		}

		right++
	}

	return minWindow
}
