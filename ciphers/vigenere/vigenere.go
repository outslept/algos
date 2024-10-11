package vigenere

func vigenereCipher(text, key string) string {
	result := ""
	keyIndex := 0
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			shift := int(key[keyIndex%len(key)] - 'A')
			result += string((char-'A'+rune(shift))%26 + 'A')
			keyIndex++
		} else if char >= 'a' && char <= 'z' {
			shift := int(key[keyIndex%len(key)] - 'a')
			result += string((char-'a'+rune(shift))%26 + 'a')
			keyIndex++
		} else {
			result += string(char)
		}
	}
	return result
}
