package kata

import (
	"fmt"
	"strings"
)

func CountPoints(games []string) int {
	var total int
	for _, game := range games {
		x, y := int(game[0]), int(game[2])
		switch {
		case x > y:
			total += 3
		case x == y:
			total += 1
		}
	}
	return total
}

func CountPoints2(games []string) int {
	var total int
	for _, game := range games {
		score := strings.Split(game, ":")
		x, y := score[0], score[1]
		switch {
		case x > y:
			total += 3
		case x == y:
			total += 1
		}
	}
	return total
}

func CountPoints3(games []string) int {
	var x, y, total int
	for _, game := range games {
		fmt.Sscanf(game, "%d:%d", &x, &y)
		switch {
		case x > y:
			total += 3
		case x == y:
			total += 1
		}
	}
	return total
}
