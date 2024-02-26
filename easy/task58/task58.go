package easy

// 58. Length of Last Word

func LengthOfLastWord(s string) int {
	words := customSplit(s)

	return len(words[len(words)-1])
}

func customSplit(s string) []string {
	if s == "" {
		return make([]string, 0)
	}

	words := make([]string, 0)
	prevIndex := 0
	for i := 0; i < len(s); i++ {

		if (s[i]) == ' ' {
			spaces := countSpaces(s, i)
			if i > prevIndex {
				words = append(words, s[prevIndex:i])
			}
			prevIndex = i + spaces
			i = prevIndex
		}
	}
	if s[len(s)-1] != ' ' {
		words = append(words, s[prevIndex:])
	}
	return words
}

func countSpaces(s string, startIndex int) int {
	var res int
	for i := startIndex; i < len(s) && s[i] == ' '; i++ {
		res++
	}
	return res
}
