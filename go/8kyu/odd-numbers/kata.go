package kata

func OddCount(n int) int{
	if n == 1 {
		return n
	}
	return int(n / 2)
}