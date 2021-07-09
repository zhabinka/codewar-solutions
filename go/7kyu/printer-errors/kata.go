package kata

import "fmt"

func PrinterError(s string) string {
	var errorsCount int
	for _, value := range s {
		if value > 'm' {
			errorsCount++
		}
	}
	output := fmt.Sprintf("%d/%d", errorsCount, len(s))

	return output
}
