package atbash

func atbashCipher(text string) string {
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string('Z' - (char - 'A'))
		} else if char >= 'a' && char <= 'z' {
			result += string('z' - (char - 'a'))
		} else {
			result += string(char)
		}
	}
	return result
}
