package kata

func CountMonkeys(n int) []int {
	var monkeysCount []int
	for i := 1; i <= n; i++ {
		monkeysCount = append(monkeysCount, i)
	}
	return monkeysCount
}

func CountMonkeys2(n int) []int {
	monkeysCount := make([]int, n)
	for i := range monkeysCount {
		monkeysCount[i] = i + 1
	}
	return monkeysCount
}
