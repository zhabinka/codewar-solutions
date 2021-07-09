package kata

import "fmt"

func Divisors(n int) int {
	if n == 1 {
		return 1
	}

	var counter int
	for i := 1; i <= n/2; i++ {
		fmt.Println(n, i, n%i, n%i == 0)
		if n%i == 0 {
			counter++
		}
	}

	return counter + 1
}
