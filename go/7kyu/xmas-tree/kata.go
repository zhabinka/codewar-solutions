package kata

import "strings"

func makeLayer(underscoresCount, latticesCount int) string {
	scip := strings.Repeat("_", underscoresCount)
	branch := strings.Repeat("#", latticesCount)
	layer := scip + branch + scip

	return layer
}

func XMasTree(height int) []string {
	XMasTree := []string{}

	if height == 0 {
		return XMasTree
	}

	length := height*2 - 1

	for i := 1; i <= length; i += 2 {
		layer := makeLayer((length-i)/2, i)
		XMasTree = append(XMasTree, layer)
	}

	trunk := makeLayer((length-1)/2, 1)
	XMasTree = append(XMasTree, trunk, trunk)

	return XMasTree
}
