package boyermoorehorspool

func boyerMooreHorspool(text, pattern string) []int {
	n, m := len(text), len(pattern)
	if m > n {
		return []int{}
	}

	skip := make(map[byte]int)
	for i := 0; i < 256; i++ {
		skip[byte(i)] = m
	}
	for i := 0; i < m-1; i++ {
		skip[pattern[i]] = m - 1 - i
	}

	var result []int
	i := 0
	for i <= n-m {
		j := m - 1
		for j >= 0 && pattern[j] == text[i+j] {
			j--
		}
		if j < 0 {
			result = append(result, i)
			i++
		} else {
			i += skip[text[i+m-1]]
		}
	}
	return result
}
