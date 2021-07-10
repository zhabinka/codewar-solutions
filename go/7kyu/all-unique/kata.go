package kata

func HasUniqueChar(str string) bool {
	charIn := make(map[rune]bool)

	for _, char := range str {
		if charIn[char] {
			return false
		}
		charIn[char] = true
	}

	return true
}
