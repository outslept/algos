package lipogram

import (
	"strings"
	"unicode"
)

func isLipogram(s string, excludedLetter rune) bool {
	s = strings.ToLower(s)
	excludedLetter = unicode.ToLower(excludedLetter)
	
	for _, char := range s {
			if char == excludedLetter {
					return false
			}
	}
	
	return true
}
