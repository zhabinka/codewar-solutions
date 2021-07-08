package kata

import "math"

func SumCubes(n int) int {
	var acc int
	for i := 1; i <= n; i++ {
		acc += int(math.Pow(float64(i), 3))
	}

	return acc
}
