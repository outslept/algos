package pangram

import (
	"strings"
	"unicode"
)

func isPangram(s string) bool {
	seen := make(map[rune]bool)
	s = strings.ToLower(s)

	for _, char := range s {
			if unicode.IsLetter(char) {
					seen[char] = true
			}
	}

	return len(seen) == 26
}
