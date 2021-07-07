package kata

func ExpressionMatter(a int, b int, c int) int {
	var max int
	results := []int{
		a * b * c,
		(a * b) + c,
		a + (b * c),
		a + b + c,
		(a + b) * c,
		a * (b + c),
	}
	for _, result := range results {
		if result > max {
			max = result
		}
	}
	return max
}
