package main

import "fmt"

// ANSI color codes
const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

// ValidateAndPrintSelector checks a selector's validation result and prints
// the message in green (valid) or red (invalid), appending "(pass)" or "(failed)".
func ValidateAndPrintSelector(komunikat string, valid bool, selector string) {
	color := Green
	status := "(pass)" // Default status for valid selectors

	if !valid {
		color = Red
		status = "(failed)"
	}

	fmt.Printf("%s: %s%s %s%s\n", komunikat, color, selector, status, Reset)
}
