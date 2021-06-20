package kata

func century(year int) int {
	var added int
	if year%100 != 0 {
		added = 1
	}
	return int(float64(year)/100) + added
}
