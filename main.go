package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Function to load the JSON file into FileContent structure
func LoadJSONFile(fileName string) (FileContent, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return FileContent{}, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	var content map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&content); err != nil {
		return FileContent{}, fmt.Errorf("Error decoding JSON: %v", err)
	}

	return FileContent{content: content}, nil
}

func main() {
	// Sample input for testing
	fileName := "test1.json"
	fileContent, err := LoadJSONFile(fileName)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	// Selector examples for testing
	selectors := []Selector{
		{file: fileContent, hook: "echo.kkk.aaa.sss[]", dataType: "string"},
		{file: fileContent, hook: "metadata.version", dataType: "number"},
		{file: fileContent, hook: "categories[].items[].id", dataType: "string"},
	}

	// Reset counters before testing
	ResetCounters()

	// Loop through selectors and test
	for _, sel := range selectors {
		valid, err := sel.CheckSelector()
		if err != nil {
			fmt.Println("Error checking selector:", err)
			continue
		}

		if valid {
			TestJSON("Test Message", "green", sel.hook)
		} else {
			TestJSON("Test Message", "red", sel.hook)
		}
	}

	// Print the test results
	PrintTestResults()
}
