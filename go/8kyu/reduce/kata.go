package kata

func Grow(arr []int) int{
	result := 1
	for _, num := range arr {
		if num > 0 {
			result *= num
		}
	}
  return result
}
