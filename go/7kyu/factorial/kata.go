package kata

func Factorial(n int) int {
	acc := 1
	for i := 1; i <= n; i++ {
		acc *= i
	}

	return acc
}

func Factorial2(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial2(n-1)
}
