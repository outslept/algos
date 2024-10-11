package vernam

func vernamCipher(text, key string) string {
	result := ""
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key[i%len(key)])
	}
	return result
}
