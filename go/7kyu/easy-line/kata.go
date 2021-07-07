package kata

import (
	"math/big"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n - 1)
}

func EasyLine(n int) string {
	if n == 0 {
		return "1"
	}
	var fact big.Int
	var fact2 big.Int
	f1 := fact.MulRange(1, int64(n * 2))
	f2 := fact2.MulRange(1, int64(n))
	f3 := f2.Mul(f2, f2)
	f4 := f3.Div(f1, f3)

	return f4.String()
}
