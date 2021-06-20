package kata

func average(a, b, c int) int {
	return (a + b + c) / 3
}

func GetGrade(a, b, c int) rune {
	score := average(a, b, c)
	switch {
	case score <= 100 && score >= 90:
		return 'A'
	case score < 90 && score >= 80:
		return 'B'
	case score < 80 && score >= 70:
		return 'C'
	case score < 70 && score >= 60:
		return 'D'
	default:
		return 'F'
	}
}
