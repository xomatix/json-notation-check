package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileType string

const (
	JSON FileType = ".json"
	JML  FileType = ".jml"
)

// CheckFile takes a filename and a filetype and prints an error message if the file
// does not match the given filetype or if it is empty.
func CheckFile(fileName string, fileType FileType) error {

	if len(fileName) < 5 || fileName[len(fileName)-len(fileType):] != string(fileType) {
		return fmt.Errorf("input file needs to be %s", fileType)
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		return fmt.Errorf("input file is empty")
	}
	return nil
}

func OpenAndReadFile(fileName string) string {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(fileContent)
}

func FillSelectors(content string, sel []Selector) []Selector {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(content), &m)

	newSels := make([]Selector, len(sel))

	for i, v := range sel {
		newSels[i] = Selector{file: FileContent{content: m}, hook: v.hook, dataType: v.dataType}
	}
	return newSels
}
