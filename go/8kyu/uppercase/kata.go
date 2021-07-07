package kata

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, value := range s {
		char := string(value)
		if char != strings.ToUpper(char) {
			return false
		}
	}
	return true
}