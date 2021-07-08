package kata

import (
	"fmt"
	"strings"
)

func PartList(list []string) string {
	var variants string

	for i := 1; i < len(list); i++ {
		partInit := list[:i]
		partFinal := list[i:]
		variants += fmt.Sprintf(
			"(%s, %s)",
			strings.Join(partInit, " "),
			strings.Join(partFinal, " "),
		)
	}

	return variants
}

func PartList2(list []string) string {
	var sb strings.Builder

	for i := 1; i < len(list); i++ {
		partInit := list[:i]
		partFinal := list[i:]
		fmt.Fprintf(
			&sb,
			"(%s, %s)",
			strings.Join(partInit, " "),
			strings.Join(partFinal, " "),
		)
	}

	return sb.String()
}
