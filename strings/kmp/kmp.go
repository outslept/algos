package kmp

func kmp(text, pattern string) []int {
	n, m := len(text), len(pattern)
	if m > n {
		return []int{}
	}

	pi := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}

	var result []int
	j = 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			result = append(result, i-m+1)
			j = pi[j-1]
		}
	}
	return result
}
