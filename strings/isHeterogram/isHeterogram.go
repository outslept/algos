package heterogram

import "unicode"

func isHeterogram(s string) bool {
	seen := make(map[rune]bool)

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
