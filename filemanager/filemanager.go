package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// Reads line of text from file and returns that text.
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	// The "defer" keyword ensures a method only runs when the surrounding function
	// is done. Saves us trouble of manually writing file.Close() in multiple places.

	// Only close file if function executes
	defer file.Close()

	// The bufio package provides utility functions for dealing w/ input and output data
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close() Comment out manual file closes
		return nil, errors.New("Failed to read line in file.")
	}

	// file.Close()
	return lines, nil
}

// Create JSON file and write data to it.
func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	// Only close file if function executes
	defer file.Close()

	// JSON package, NewEncoder() that converts data into JSON format
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Failed to convert data to JSON.")
	}

	// file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
