package kata

func SumSquare(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number * number
	}
	return sum
}
