package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"golang.org/x/sys/windows"
)

// Color codes for console output
const (
	greenColor  = "\x1b[32m"
	redColor    = "\x1b[31m"
	yellowColor = "\x1b[33m"
	resetColor  = "\x1b[0m"
)

// Status represents the result of interpretation
type Status int

const (
	Warning Status = iota
	Success
)

// DataTypeAnalyzer handles JSON interpretation
type DataTypeAnalyzer struct {
	input     string
	selectors []string
	result    map[string]string
	expected  map[string]string
	passed    int
	failed    int
}

// NewDataTypeAnalyzer creates a new analyzer instance
func NewDataTypeAnalyzer(input string) *DataTypeAnalyzer {
	return &DataTypeAnalyzer{
		input:     input,
		selectors: []string{},
		result:    make(map[string]string),
		expected:  make(map[string]string),
		passed:    0,
		failed:    0,
	}
}

// enableWindowsANSIColors enables ANSI color support on Windows
func enableWindowsANSIColors() {
	stdout := windows.Handle(os.Stdout.Fd())
	var mode uint32
	windows.GetConsoleMode(stdout, &mode)
	windows.SetConsoleMode(stdout, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

// AddSelector adds a specific selector to analyze with expected type
func (a *DataTypeAnalyzer) AddSelector(selector string, expectedType string) {
	a.selectors = append(a.selectors, selector)
	a.expected[selector] = expectedType
}

// Announce creates a descriptive message about the analysis
func (a *DataTypeAnalyzer) Announce() string {
	return fmt.Sprintf("Analyzing JSON data with %d selectors", len(a.selectors))
}

// Interpret processes the JSON and determines data types
func (a *DataTypeAnalyzer) Interpret() (Status, error) {
	// Parse the JSON
	var data map[string]interface{}
	err := json.Unmarshal([]byte(a.input), &data)
	if err != nil {
		return Warning, fmt.Errorf("invalid JSON: %v", err)
	}

	// If no selectors specified, analyze all top-level keys
	if len(a.selectors) == 0 {
		for key, value := range data {
			a.result[key] = a.determineType(value)
		}
	} else {
		// Analyze specific selectors
		for _, selector := range a.selectors {
			parts := strings.Split(selector, ".")
			if len(parts) < 2 {
				a.result[selector] = "invalid selector"
				continue
			}

			obj := data[parts[0]]
			if objMap, ok := obj.(map[string]interface{}); ok {
				value := objMap[parts[1]]
				a.result[selector] = a.determineType(value)
			}
		}
	}

	return Success, nil
}

// determineType identifies the specific type of a JSON value
func (a *DataTypeAnalyzer) determineType(value interface{}) string {
	if value == nil {
		return "null"
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return "string"
	case reflect.Float64:
		// JSON numbers are always float64 in Go
		// Check if it's a whole number
		if float64(int(value.(float64))) == value.(float64) {
			return "int"
		}
		return "float"
	case reflect.Bool:
		return "bool"
	case reflect.Slice:
		// Check if the slice is empty
		slice := reflect.ValueOf(value)
		if slice.Len() == 0 {
			return "empty_array"
		}
		return "array"
	case reflect.Map:
		return "object"
	default:
		return "unknown"
	}
}

// PrintResults displays the interpretation results with color-coded output
func (a *DataTypeAnalyzer) PrintResults() {
	fmt.Println("Interpretation Results:")
	for selector, dataType := range a.result {
		// Check if the detected type matches the expected type
		expectedType, hasExpectation := a.expected[selector]

		if hasExpectation {
			if dataType == expectedType ||
				(expectedType == "array" && dataType == "empty_array") {
				// Pass cases: exact type match or empty array when array is expected
				color := greenColor
				if dataType == "empty_array" {
					color = yellowColor
				}

				fmt.Printf("%s: %s%s%s %s%s%s\n",
					selector,
					color, dataType, resetColor,
					color, "✓", resetColor)
				a.passed++
			} else {
				fmt.Printf("%s: %s%s%s %s%s%s (Expected: %s)\n",
					selector,
					redColor, dataType, resetColor,
					redColor, "✗", resetColor,
					expectedType)
				a.failed++
			}
		} else {
			// If no expectation was set, just print the type
			fmt.Printf("%s: %s\n", selector, dataType)
		}
	}
}

// PrintSummary displays a summary of passed and failed tests
func (a *DataTypeAnalyzer) PrintSummary() {
	fmt.Println("\nTest Summary:")
	fmt.Printf("%sPassed: %d%s\n", greenColor, a.passed, resetColor)
	fmt.Printf("%sFailed: %d%s\n", redColor, a.failed, resetColor)
	fmt.Printf("Total Tests: %d\n", a.passed+a.failed)
}

func main() {
	// Enable ANSI colors on Windows
	enableWindowsANSIColors()

	// Example JSON input with type mismatches and empty array
	jsonInput := `{
		"person": {
			"name": 42,
			"age": 30,
			"isStudent": true,
			"grades": [],
			"items": []
		}
	}`

	// Create analyzer
	analyzer := NewDataTypeAnalyzer(jsonInput)

	// Add specific selectors to analyze with expected types
	analyzer.AddSelector("person.name", "string")    // This will fail
	analyzer.AddSelector("person.age", "int")        // This will pass
	analyzer.AddSelector("person.isStudent", "bool") // This will pass
	analyzer.AddSelector("person.grades", "array")   // This will pass with yellow
	analyzer.AddSelector("person.items", "array")    // This will pass with yellow

	// Print announcement
	fmt.Println(analyzer.Announce())

	// Interpret the JSON
	status, err := analyzer.Interpret()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print results based on status
	switch status {
	case Success:
		analyzer.PrintResults()
		analyzer.PrintSummary()
	case Warning:
		fmt.Println("Warning: Partial or incomplete analysis")
	}
}
