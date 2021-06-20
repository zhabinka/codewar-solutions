package kata

import "strings"

func reverse(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

func IsPalindrome(str string) bool {
	return strings.ToLower(str) == strings.ToLower(reverse(str))
	//return strings.EqualFold(str, reverse(str))
}

func reverse2(str string) string {
	var result string
	for _, letter := range str {
		result = string(letter) + result
	}
	return result
}
