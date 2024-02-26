package task58

import "strings"

func LengthOfLastWord_Fields(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	return len(words[len(words)-1])
}
