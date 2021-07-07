package kata

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	set := strings.Split(in, " ")
	base, err := strconv.Atoi(set[0])
	if err != nil {
		fmt.Errorf("value cannot be cast to type number")
	}
	min, max := base, base

	for _, v := range set {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Errorf("value cannot be cast to type number")
		}
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}

func HighAndLow2(in string) string {
	set := strings.Fields(in)
	length := len(set)
	numbers := make([]int, 0, length)

	for _, v := range set {
		num, _ := strconv.Atoi(v)
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	return fmt.Sprintf("%d %d", numbers[length-1], numbers[0])
}
