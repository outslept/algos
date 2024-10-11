package substitution

func substitutionCipher(text string, key map[rune]rune) string {
	result := ""
	for _, char := range text {
		if newChar, ok := key[char]; ok {
			result += string(newChar)
		} else {
			result += string(char)
		}
	}
	return result
}
