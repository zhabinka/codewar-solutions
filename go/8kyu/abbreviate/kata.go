package kata

import (
	"strings"
)

func AbbrevName(name string) string {
	list := strings.Split(name, " ")
	var abbrev []string
	for _, v := range list {
		abbrev = append(abbrev, strings.ToUpper(string(v[0])))
	}
	return strings.Join(abbrev, ".")
}
