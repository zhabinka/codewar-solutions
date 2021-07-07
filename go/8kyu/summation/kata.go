package kata

func Summation(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func Summation2(n int) int {
	return n * (n + 1) / 2
}

func Summation3(n int) int {
	if n < 1 {
		return 0
	}
	return n + Summation3(n-1)
}

//func Summation4(n int) int {
//	iter := func(acc, i int) int {
//		if i > n {
//			return acc
//		}
//		return iter(acc, i+1)
//	}
//	return iter(0, 1)
//}