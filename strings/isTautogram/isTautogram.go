package tautogram

import "strings"

func isTautogram(s string) bool {
    words := strings.Fields(s)
    if len(words) < 2 {
        return true
    }
    
    firstLetter := strings.ToLower(string(words[0][0]))
    
    for _, word := range words[1:] {
        if strings.ToLower(string(word[0])) != firstLetter {
            return false
        }
    }
    
    return true
}
