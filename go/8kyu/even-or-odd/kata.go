package kata

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func EvenOrOdd2(number int) string {
	return []string{"Even", "Odd"}[number%2]
}

func EvenOrOdd3(number int) string {
	return []string{"Even", "Odd"}[number&1]
}
