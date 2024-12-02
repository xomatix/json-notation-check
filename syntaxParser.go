package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseCommand(input string) {
	// Define the regex pattern
	// pattern := `^\$(\w+(\[\]){0,1}\s*.{0,1})*:\w+$`
	pattern := `\$(\w+(\[\]){0,1}\s*\.{0,1}){1,}:\w+`

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Find all matches in the input string
	matches := re.FindAllString(input, -1)
	// fmt.Printf("%v", matches)
	// Display the matched strings in the array
	fmt.Println("Commands Strings:\n")
	for i, match := range matches {
		fmt.Println("Command ", i+1)
		fmt.Println("=============================")
		// match = strings.ReplaceAll(match, " ", "")
		// match = strings.ReplaceAll(match, "\n", "")
		lines := strings.Split(match, ":")

		fmt.Println("selector")
		fmt.Println(strings.Trim(lines[0], "$"))

		fmt.Println("\ntype")
		fmt.Println(lines[1])

		fmt.Println("=============================\n")
	}
}
