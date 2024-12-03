package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("You need to provide checker file and an input file with .json extension to run this program")
		return
	}

	checkerFile := os.Args[1]
	err := CheckFile(checkerFile, JML)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	inputFile := os.Args[2]
	err = CheckFile(inputFile, JSON)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// fmt.Printf("File checker: %s\n", checkerFile)
	// fmt.Printf("Input json: %s\n", inputFile)

	checkerFileContent := OpenAndReadFile(checkerFile)
	inputFileContent := OpenAndReadFile(inputFile)

	selArr := ParseCommand(checkerFileContent)

	selArr = FillSelectors(inputFileContent, selArr)

	for _, sel := range selArr {
		komunikat := "Validation Result"

		ok, _ := sel.CheckSelector()

		// Use validation result to set color
		ValidateAndPrintSelector(komunikat, ok, sel.hook)

		if ok {
			fmt.Printf("DataType: %s\n", sel.dataType)
		} else {
			fmt.Printf("DataType: %s\n", sel.dataType)
			fmt.Printf("Error: %s\n", "Implement error")
		}

		fmt.Printf("\n\n")
	}

	// fmt.Printf("Commands: %v\n", commands)
	// fmt.Printf("Command args: %v\n", commandArgs)

	// println("Hello, World!")
}
