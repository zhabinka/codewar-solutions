package kata

import "math"

func combat(health, damage float64) float64 {
	if health > damage {
		return health - damage
	}
	return 0.0
}

func combat2(health, damage float64) float64 {
	return math.Max(health-damage, 0)
}
