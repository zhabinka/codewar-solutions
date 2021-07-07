package kata

import "math"

func SquareOrSquareRoot2(arr []int) []int {
	result := make([]int, len(arr))
	for i, number := range arr {
		sqrt := math.Sqrt(float64(number))
		if math.Ceil(sqrt) != sqrt {
			result[i] = number * number
		} else {
			result[i] = int(sqrt)
		}
	}
	return result
}

func SquareOrSquareRoot(arr []int) []int {
	result := make([]int, len(arr))
	for i, num := range arr {
		sqrt, fractional := math.Modf(math.Sqrt(float64(num)))
		if fractional == 0 {
			result[i] = int(sqrt)
		} else {
			result[i] = num * num
		}
	}
	return result
}
