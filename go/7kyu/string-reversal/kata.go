package kata

import (
	"strings"
)

func Solve(s string) string {
	reversal := make([]string, 0, len(s))
	var ii int

	for i := len(s) - 1; i >= 0 ; i-- {
		current := string(s[i])
		if current == " " {
			continue
		}
		if string(s[ii]) == " " {
			reversal = append(reversal, " ")
		  ii++
		}
		ii++

		reversal = append(reversal, current)
	}

	return strings.Join(reversal, "")
}
