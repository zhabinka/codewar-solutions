package kata

func DoubleValues(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, number := range numbers {
		result[i] = number*2
	}
	return result
}

func DoubleValues2(numbers []int) (result []int) {
	result = make([]int, len(numbers))
	for i, number := range numbers {
		result[i] = number*2
	}
	return
}
