package task2108

func firstPalindrome(words []string) string {
	for _, word := range words {
		if word == reverse(word) {
			return word
		}
	}
	return ""
}

func reverse(s string) string {
	length := len(s)
	runes := make([]rune, length)

	for _, runeValue := range s {
		length--
		runes[length] = runeValue
	}

	return string(runes[length:])
}
