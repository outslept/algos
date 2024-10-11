package isogram

import (
	"strings"
	"unicode"
)

func isIsogram(s string) bool {
	seen := make(map[rune]bool)
	s = strings.ToLower(s)

	for _, char := range s {
		if unicode.IsLetter(char) {
			if seen[char] {
				return false
			}
			seen[char] = true
		}
	}

	return true
}
