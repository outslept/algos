package rabinkarp

func rabinKarp(text, pattern string) []int {
	n, m := len(text), len(pattern)
	if m > n {
		return []int{}
	}

	const prime = 101
	var powM = 1
	for i := 0; i < m-1; i++ {
		powM = (powM * 256) % prime
	}

	var textHash, patternHash int
	for i := 0; i < m; i++ {
		patternHash = (patternHash*256 + int(pattern[i])) % prime
		textHash = (textHash*256 + int(text[i])) % prime
	}

	var result []int
	for i := 0; i <= n-m; i++ {
		if patternHash == textHash {
			if text[i:i+m] == pattern {
				result = append(result, i)
			}
		}
		if i < n-m {
			textHash = (256*(textHash-int(text[i])*powM) + int(text[i+m])) % prime
			if textHash < 0 {
				textHash += prime
			}
		}
	}
	return result
}
