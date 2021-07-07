package kata

import "strings"

func RepeatStr(times int, value string) string {
	var result string
	for i := 0; i < times; i++ {
		result += value
	}
	return result
}

func RepeatStr2(times int, value string) string {
	return strings.Repeat(value, times)
}

func RepeatStr3(times int, value string) string {
	var result strings.Builder
	for i := 0; i < times; i++ {
		result.WriteString(value)
	}
	return result.String()
}
