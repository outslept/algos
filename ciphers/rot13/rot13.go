package rot13

func rot13(text string) string {
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+13)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result += string((char-'a'+13)%26 + 'a')
		} else {
			result += string(char)
		}
	}
	return result
}
