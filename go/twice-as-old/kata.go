package kata

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	diff := dadYearsOld - sonYearsOld*2
	return int(math.Abs(float64(diff)))
}
