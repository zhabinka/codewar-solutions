package kata

func multipleOfIndex(numbers []int) []int {
	result := []int{}
	for i := 1; i < len(numbers); i++ {
		num := numbers[i]
		if num%(i) == 0 {
			result = append(result, num)
		}
	}
	return result
}
