package kata

import "strings"

func FindShort(s string) int {
	words := strings.Split(s, " ")
	minLength := len(words[0])

	for _, word := range words {
		length := len(word)
		if length < minLength {
			minLength = length
		}
	}

	return minLength
}
