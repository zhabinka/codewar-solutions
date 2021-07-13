package kata

func include(nums []int, n int) bool {
	for _, num := range nums {
		if num == n {
			return true
		}
	}

	return false
}

func Solve(arr []int) int {
	for _, num := range arr {
		if !include(arr, -num) {
			return num
		}
	}

	return 0
}
