package kata

import "strings"

func solution(str, ending string) bool {
	index := len(str) - len(ending)
	if index < 0 {
		return false
	}

	tail := str[index:]
	return tail == ending
}

func solution2(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
