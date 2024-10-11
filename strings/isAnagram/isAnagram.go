package anagram

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
			return false
	}
	
	count := make(map[rune]int)
	
	for _, char := range s1 {
			count[char]++
	}
	
	for _, char := range s2 {
			count[char]--
			if count[char] < 0 {
					return false
			}
	}
	
	return true
}
