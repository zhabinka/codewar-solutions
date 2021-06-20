package kata

import (
	"fmt"
	"strings"
)

func countSheep(num int) string {
	var builder strings.Builder
	for i := 1; i <= num; i++ {
		fmt.Fprintf(&builder, "%d sheep...", i)
	}
	return builder.String()
}

func countSheep2(num int) string {
	var result string
	for i := 1; i <= num; i++ {
		result += fmt.Sprintf("%d sheep...", i)
	}
	return result
}
