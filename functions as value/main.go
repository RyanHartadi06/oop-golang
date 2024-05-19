package main

import (
	"fmt"
	"strings"
)

func numberTotal(a, b int) int {
	return a + b
}

// functions as params
type Formatter func(string) string

func formatCase(txt string, formatter Formatter) {
	formatted := formatter(txt)
	fmt.Println(formatted)
}

func upperCase(txt string) string {
	return strings.ToUpper(txt)
}

func toLowerCase(txt string) string {
	return strings.ToLower(txt)
}

func main() {
	total := numberTotal
	fmt.Println(total(1, 2))

	toUpper := upperCase
	formatCase("hello world", toUpper)
}
