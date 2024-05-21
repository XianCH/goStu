package main

import (
	"fmt"
	"io/ioutil"
	"os"

	errortest "github.com/x14n/goStu/basic/error"
)

// errortest.TestErrorChain()
// errortest.TestErrorJoin()
// errortest.TestRootCaue()
// errortest.TestContextError()
// readFile reads the content of a file and returns it as a string.
func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errortest.Wrap(err, fmt.Sprintf("failed to read file: %v", err))
	}
	return string(data), nil
}

// processFile processes the content of a file.
func processFile(filename string) error {
	content, err := readFile(filename)
	if err != nil {
		return errortest.Wrap(err, "failed to process file")
	}

	fmt.Println("File content:", content)
	return nil
}

func main() {
	filename := "nonexistent.txt"
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	err := processFile(filename)
	// _, err := readFile(filename)
	if err != nil {
		errortest.ErrorHelper(err)
		// fmt.Println(err)
	}
}
