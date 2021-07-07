package kata

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	// Конкатенация очень дорогая операция
	return "Hello, " + name  + " how are you doing today?"
}

func Greet2(name string) string {
	// Print делает рефлексию, анализирует структуру кода
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func Greet3(name string) string {
	var builder strings.Builder
	builder.WriteString("Hello, ")
	builder.WriteString(name)
	builder.WriteString(" how are you doing today?")
	return builder.String()
}


/*
1. Пишем тест.


 */