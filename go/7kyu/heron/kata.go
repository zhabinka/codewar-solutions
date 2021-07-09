package kata

import (
	"math"
)

func Heron(a, b, c int) float32 {
	he := float64(a+b+c) / 2
	area := math.Sqrt(he * (he - float64(a)) * (he - float64(b)) * (he - float64(c)))
	return float32(math.Round(area*100) / 100)
}
