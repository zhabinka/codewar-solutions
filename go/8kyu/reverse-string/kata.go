package kata

import "strings"

func Reverse(str string) string {
	var result string
	for _, char := range str {
		result = string(char) + result
	}
	return result
}

func Reverse2(word string) string {
	var b strings.Builder
	for i := len(word) - 1; i > -1; i-- {
		b.WriteByte(word[i])
	}
	return b.String()
}
