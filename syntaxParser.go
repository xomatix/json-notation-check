package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseCommand(input string) []Selector {
	// Define the regex pattern
	// pattern := `^\$(\w+(\[\]){0,1}\s*.{0,1})*:\w+$`
	pattern := `\$(\w+(\[\]){0,1}\s*\.{0,1}){1,}:\w+`

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return make([]Selector, 0)
	}

	// Find all matches in the input string
	matches := re.FindAllString(input, -1)
	// fmt.Printf("%v", matches)
	// Display the matched strings in the array
	fmt.Println("Commands Strings:\n")
	selectors := make([]Selector, len(matches))
	for i, match := range matches {
		// fmt.Println("Command ", i+1)
		// fmt.Println("=============================")
		// // match = strings.ReplaceAll(match, " ", "")
		// match = strings.ReplaceAll(match, "\n", "")
		lines := strings.Split(match, ":")

		// fmt.Println("selector")
		// fmt.Println(strings.Trim(lines[0], "$"))

		// fmt.Println("\ntype")
		// fmt.Println(lines[1])

		// fmt.Println("=============================\n")
		selectors[i] = Selector{hook: strings.Trim(lines[0], "$"), dataType: lines[1]}
	}

	return selectors
}

type FileContent struct {
	content map[string]interface{}
}

type Selector struct {
	file     FileContent
	hook     string
	dataType string
}

func (s Selector) CheckSelector() (bool, error) {
	//json input file
	hooks := strings.Split(s.hook, ".")

	return checkSymbol(hooks, s.file.content, s.dataType), nil //fmt.Errorf("Not implemented!")
}

func checkSymbol(selector []string, m map[string]interface{}, dataType string) bool {
	hooks := selector
	if len(hooks) == 1 {
		switch dataType {
		case "string":
			return checkForString(hooks[0], m)
		case "number":
			return checkForNumber(hooks[0], m)
		case "bool":
			return checkForBool(hooks[0], m)
		case "array":
			return checkForArray(hooks[0], m)
		case "object":
			return checkForObject(hooks[0], m)
		// case "null":
		// 	return checkForNull(hooks[0], m)
		default:
			return false
		}
	} else if strings.Contains(hooks[0], "[]") {
		formated := strings.ReplaceAll(hooks[0], "[]", "")
		sel, ok := m[formated].([]interface{})

		if !ok {
			return false
		}
		for _, v := range sel {
			res := checkSymbol(hooks[1:], v.(map[string]interface{}), dataType)
			if !res {
				return false
			}
		}
		return true
	} else {
		sel, ok := m[hooks[0]].(map[string]interface{})
		if !ok {
			return false
		}

		return checkSymbol(hooks[1:], sel, dataType)
	}
}

func checkForString(selector string, m map[string]interface{}) bool {
	if _, ok := m[selector].(string); ok {
		return true
	}
	return false
}

func checkForNumber(selector string, m map[string]interface{}) bool {
	if _, ok := m[selector].(float64); ok {
		return true
	}
	return false
}

func checkForBool(selector string, m map[string]interface{}) bool {
	if _, ok := m[selector].(bool); ok {
		return true
	}
	return false
}

func checkForArray(selector string, m map[string]interface{}) bool {
	if _, ok := m[selector].([]interface{}); ok {
		return true
	}
	return false
}

func checkForObject(selector string, m map[string]interface{}) bool {
	if _, ok := m[selector].(map[string]interface{}); ok {
		return true
	}
	return false
}

// func checkForNull(selector string, m map[string]interface{}) bool {
// 	if _, ok := m[selector].(); ok {
// 		return true
// 	}
// 	return false
// }
