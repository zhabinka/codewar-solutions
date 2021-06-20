package kata

import "fmt"

func BonusTime(salary int, bonus bool) string {
	if bonus {
		salary *= 10
	}
	return fmt.Sprintf("\u00A3%d", salary)
}
