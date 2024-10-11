package beaufort

func beaufortCipher(text, key string) string {
	result := ""
	keyIndex := 0
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			shift := int(key[keyIndex%len(key)] - 'A')
			result += string((26+shift-int(char-'A'))%26 + 'A')
			keyIndex++
		} else if char >= 'a' && char <= 'z' {
			shift := int(key[keyIndex%len(key)] - 'a')
			result += string((26+shift-int(char-'a'))%26 + 'a')
			keyIndex++
		} else {
			result += string(char)
		}
	}
	return result
}
