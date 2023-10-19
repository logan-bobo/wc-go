package cmd

import (
	"os"
	"strings"
)

func CountChars(fileName string) int {
	var charCount int

    bytesRead, _ := os.ReadFile(fileName)
    fileContent := string(bytesRead)
    lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		arrLine := strings.Split(line, "")
		for i := 0; i < len(arrLine); i++ {
			charCount++
		}

		// New line counts as one cah so add it here
		charCount++
	}

	// Last line will not contain a newline so remove
	charCount--

	return charCount
}