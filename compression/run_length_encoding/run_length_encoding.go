package rle

import (
	"fmt"
	"strings"
)

func RLEEncode(input string) string {
	if len(input) == 0 {
			return ""
	}

	var result strings.Builder
	count := 1
	current := input[0]

	for i := 1; i < len(input); i++ {
			if input[i] == current {
					count++
			} else {
					result.WriteString(fmt.Sprintf("%d%c", count, current))
					current = input[i]
					count = 1
			}
	}

	result.WriteString(fmt.Sprintf("%d%c", count, current))

	return result.String()
}

func RLEDecode(input string) (string, error) {
	var result strings.Builder
	var count strings.Builder

	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			count.WriteByte(input[i])
		} else {
			n := 0
			if _, err := fmt.Sscanf(count.String(), "%d", &n); err != nil {
				return "", fmt.Errorf("invalid input: %d", i)
			}
			result.WriteString(strings.Repeat(string(input[i]), n))
			count.Reset()
		}
	}

	return result.String(), nil
}
