package kata

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	uniqUnion := ""
	for _, char := range s1 + s2 {
		if !strings.ContainsRune(uniqUnion, char) {
			uniqUnion += string(char)
		}
	}

	list := strings.Split(uniqUnion, "")
	sort.Strings(list)

	result := ""
	for _, value := range list {
		result += value
	}
	return result
}

