package main

import (
	"fmt"
)

// Counter variables for tracking passed and failed tests
var passedTests int
var failedTests int

// Reset the counters before each run
func ResetCounters() {
	passedTests = 0
	failedTests = 0
}

// TestJSON handles the test results, simulating passing and failing conditions
func TestJSON(komunikat, kolor, selector string) {
	fmt.Printf("Testing with Message: %s, Color: %s, Selector: %s\n", komunikat, kolor, selector)

	// Simulate test result (replace with actual logic)
	if selector == "echo.kkk.aaa.sss[]" {
		fmt.Printf("Test passed for selector '%s': Found the expected title.\n", selector)
		passedTests++
	} else {
		fmt.Printf("Selector %s is invalid!\n", selector)
		failedTests++
	}
}

func PrintTestResults() {
	fmt.Println("\nTest Summary:")
	fmt.Printf("Tests Passed: %d\n", passedTests)
	fmt.Printf("Tests Failed: %d\n", failedTests)

	if failedTests == 0 {
		fmt.Println("All tests passed! 🎉")
	} else {
		fmt.Println("Some tests failed. Please review the output.")
	}
}
