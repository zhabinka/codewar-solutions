package kata

import "strings"

func DNAtoRNA(dna string) string {
	rna := strings.ReplaceAll(dna, "T", "U")
	return rna
}

func DNAtoRNA2(dna string) string {
	var rna string
	for _, char := range dna {
		if string(char) == "T" {
			rna += "U"
			continue
		}
		rna += string(char)
	}
	return rna
}
