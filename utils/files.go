package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetInputFilePath() (string, error) {
	usageErr := fmt.Errorf("usage: go run %s <input file path>", os.Args[0])
	if len(os.Args) != 2 {
		return "", usageErr
	}

	path := os.Args[1]
	return path, nil
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ReadInputFile() ([]string, error) {
	path, err := GetInputFilePath()
	if err != nil {
		return nil, err
	}

	input, err := ReadFile(path)
	if err != nil {
		return nil, err
	}

	return input, err
}