package kata

import "strings"

var vowels string = "aeiou"

func Solve(s string) int {
	var max, current int
	// for i := 0; i < len(s); i++ {
	//    if strings.Contains(vowels, string(s[i])) {
	//        current++
	//    }
	//    if
	//    else {
	//        max = current
	//        current = 0
	//    }

	// }
	for _, char := range s {
		if strings.Contains(vowels, string(char)) {
			current++
		} else {
			if current > max {
				max = current
			}
			current = 0
		}
	}

	return max
}
