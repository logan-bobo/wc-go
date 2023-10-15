package cmd

import (
	"bufio"
	"os"
)

func CountLines(fileName string) int {
	var lineCount int

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}
